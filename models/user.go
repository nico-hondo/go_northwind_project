package models

type User struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Roles    string `json:"roles"`
	Token    string `json:"user_token"`
}
