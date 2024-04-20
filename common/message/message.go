package message

import "fmt"

func SystemError() string { return "システムエラー" }

func NotFound(name string) string { return fmt.Sprintf("%sが見つかりません。", name) }
func Failed(name string) string   { return fmt.Sprintf("%sに失敗しました。", name) }
