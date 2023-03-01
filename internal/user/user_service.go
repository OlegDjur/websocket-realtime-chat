package user

import (
	"context"
	"time"
)

const (
	secretKey = "secret"
)

type service struct {
	Repository
	timeeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(c context.Context, req *LoginUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)
}

func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
}
