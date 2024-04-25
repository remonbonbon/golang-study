package common

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"sync"
	"time"
)

type MyHandler struct {
	opts  MyHandlerOptions
	attrs []slog.Attr // attrs if non-empty
	mu    *sync.Mutex
	out   io.Writer
}

type MyHandlerOptions struct {
	// Level reports the minimum level to log.
	// Levels with lower levels are discarded.
	// If nil, the Handler uses [slog.LevelInfo].
	Level slog.Leveler
}

func NewMyHandler(out io.Writer, opts *MyHandlerOptions) *MyHandler {
	h := &MyHandler{out: out, mu: &sync.Mutex{}}
	if opts != nil {
		h.opts = *opts
	}
	if h.opts.Level == nil {
		h.opts.Level = slog.LevelInfo
	}
	return h
}

func (h *MyHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

func (h *MyHandler) Handle(ctx context.Context, r slog.Record) error {
	fmt.Printf("")
	buf := make([]byte, 0, 1024)

	buf = fmt.Appendf(buf, "%s", r.Time.Format(time.RFC3339Nano))
	buf = fmt.Appendf(buf, " [%v] ", r.Level)
	buf = fmt.Appendf(buf, "%s", r.Message)
	buf = fmt.Appendf(buf, "%+v", h.attrs)
	buf = fmt.Appendf(buf, "\n")

	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.out.Write(buf)
	return err
}

// 未実装。使用しない
func (h *MyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return h
	}

	h2 := *h
	h2.attrs = make([]slog.Attr, len(h.attrs))
	copy(h2.attrs, h.attrs)
	h2.attrs = append(h2.attrs, attrs...)
	return &h2
}

// 未実装。使用しない
func (h *MyHandler) WithGroup(name string) slog.Handler {
	panic("WithGroup()は未実装")
	// return h
}
