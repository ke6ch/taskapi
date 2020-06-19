package model

// User ユーザ情報
type (
	User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
