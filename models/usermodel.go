package models

type User struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserReq struct {
	ID         int
	Username   string
	Email      string
	Created_at string
}

type UserLogin struct {
	UserName string
	Password string
	Token    string
}
