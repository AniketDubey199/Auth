package models

type User struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UserReq struct {
	ID       int
	Username string
	Email    string
	Phone    string
	Password string
}
