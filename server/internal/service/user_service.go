package service

import (
	"fmt"

	"github.com/sunilkkhadka/Forum/internal/dto"
	"github.com/sunilkkhadka/Forum/internal/model"
	"github.com/sunilkkhadka/Forum/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceI interface {
	RegisterUser(registerUser *dto.RegisterUserRequest) error
}

type UserService struct {
	UserRepo repository.UserRepositoryI
}

func NewUserService(userRepo repository.UserRepositoryI) UserServiceI {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (service *UserService) RegisterUser(registerUser *dto.RegisterUserRequest) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("Couldn't hash password: %w", err)
	}

	user := &model.RegisterUser{
		Email:    registerUser.Email,
		Password: string(hashedPassword),
	}

	err = service.UserRepo.RegisterUser(user)
	if err != nil {
		return err
	}

	return nil
}
