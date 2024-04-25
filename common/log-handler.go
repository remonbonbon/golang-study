package common

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

// 人間が読みやすいように出力するロガー
type HumanHandler struct {
	opts HumanHandlerOptions
	mu   *sync.Mutex
	out  io.Writer

	attrs []slog.Attr
}

type HumanHandlerOptions struct {
	Level slog.Leveler // 最小ログレベル
}

func NewHumanHandler(out io.Writer, opts *HumanHandlerOptions) *HumanHandler {
	h := &HumanHandler{out: out, mu: &sync.Mutex{}}
	if opts != nil {
		h.opts = *opts
	}
	if h.opts.Level == nil {
		h.opts.Level = slog.LevelDebug
	}
	return h
}

func (h *HumanHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

func (h *HumanHandler) Handle(ctx context.Context, r slog.Record) error {
	// ログレベルに応じた色付け
	color := ColorRed
	switch r.Level {
	case slog.LevelDebug:
		color = ColorCyan
	case slog.LevelInfo:
		color = ColorGreen
	case slog.LevelWarn:
		color = ColorYellow
	}

	// 高速化のために[]byteバッファを確保してfmt.Appendf()していく
	buf := make([]byte, 0, 1024)

	// 固定のフィールド
	buf = fmt.Appendf(buf, "%s%s [%s] %s%s",
		r.Time.Format("2006-01-02 15:04:05.000"),
		color,
		r.Level,
		r.Message,
		ColorReset)

	// ログメソッドで追加されたフィールド
	if 0 < len(h.attrs) {
		buf = fmt.Appendf(buf, " ")
	}
	for index, value := range h.attrs {
		if 0 < index {
			buf = fmt.Appendf(buf, " ")
		}
		buf = fmt.Appendf(buf, toString(value, ""))
	}

	// With()で追加されたフィールド
	if 0 < r.NumAttrs() {
		buf = fmt.Appendf(buf, " ")
	}
	i := 0
	r.Attrs(func(a slog.Attr) bool {
		if 0 < i {
			buf = fmt.Appendf(buf, " ")
		}
		buf = fmt.Appendf(buf, toString(a, ""))
		i++
		return true
	})

	// 呼び出し元。カレントディレクトリからの相対パスにする
	_, file, line, ok := runtime.Caller(3)
	if ok {
		dir, _ := os.Getwd()
		rel, _ := filepath.Rel(dir, file)
		rel = filepath.ToSlash(rel)
		buf = fmt.Appendf(buf, " %s<%s:%d>%s", ColorBlack, rel, line, ColorReset)
	}

	buf = fmt.Appendf(buf, "\n")

	// 出力
	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.out.Write(buf)
	return err
}

// slog.Attrを文字列化
func toString(a slog.Attr, parent string) string {
	switch a.Value.Kind() {
	case slog.KindGroup:
		s := ""
		for index, value := range a.Value.Group() {
			if 0 < index {
				s += " "
			}
			name := parent + a.Key + "."
			s += toString(value, name)
		}
		return s
	default:
		return fmt.Sprintf("%s%s%s%s=%+v",
			ColorMagenta,
			parent,
			a.Key,
			ColorReset,
			a.Value.String())
	}
}

func (h *HumanHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return h
	}

	// attrsを連結した新しいインスタンスを作成
	h2 := *h
	h2.attrs = make([]slog.Attr, len(h.attrs))
	copy(h2.attrs, h.attrs)
	h2.attrs = append(h2.attrs, attrs...)
	return &h2
}

func (h *HumanHandler) WithGroup(name string) slog.Handler {
	panic("未実装")
}
