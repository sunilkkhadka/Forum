package service

import "github.com/sunilkkhadka/Forum/internal/repository"

type UserServiceI interface{}

type UserService struct {
	UserRepo repository.UserRepositoryI
}

func NewUserService(userRepo repository.UserRepositoryI) UserServiceI {
	return &UserService{
		UserRepo: userRepo,
	}
}
