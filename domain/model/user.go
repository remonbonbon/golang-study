package model

// ユーザーモデル
type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}
