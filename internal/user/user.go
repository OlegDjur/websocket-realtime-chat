package user

import "context"

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `josn:"password"`
	Email    string `json:"email"`
}

type CreateUserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRes struct {
	accessToken string
	ID          string `json:"id"`
	Username    string `json:"username"`
}

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUser(ctx context.Context, email string) (*User, error)
}

type Service interface {
	CreateUser(c context.Context, req *LoginUserReq) (*CreateUserRes, error)
	Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error)
}
