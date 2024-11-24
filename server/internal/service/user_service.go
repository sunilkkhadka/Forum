package service

import (
	"fmt"

	"github.com/sunilkkhadka/Forum/internal/dto"
	"github.com/sunilkkhadka/Forum/internal/model"
	"github.com/sunilkkhadka/Forum/internal/repository"
	"github.com/sunilkkhadka/Forum/internal/utils"
	"github.com/sunilkkhadka/Forum/internal/utils/auth"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceI interface {
	RegisterUser(registerUser *dto.BasicUserRequest) error
	LoginUser(loginUser *dto.BasicUserRequest) (*model.User, string, string, error)
}

type UserService struct {
	UserRepo repository.UserRepositoryI
}

func NewUserService(userRepo repository.UserRepositoryI) UserServiceI {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (service *UserService) RegisterUser(registerUser *dto.BasicUserRequest) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("couldn't hash password: %w", err)
	}

	user := &model.BasicUser{
		Email:    registerUser.Email,
		Password: string(hashedPassword),
	}

	err = service.UserRepo.RegisterUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (service *UserService) LoginUser(loginUser *dto.BasicUserRequest) (*model.User, string, string, error) {
	user, err := service.UserRepo.GetUserByEmail(loginUser.Email)

	if err != nil {
		return nil, "", "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginUser.Password))
	if err != nil {
		return nil, "", "", utils.ErrPasswordDoNotMatch
	}

	accessToken, refreshToken, err := auth.GenerateToken(uint(user.ID), "")
	if err != nil {
		return nil, "", "", err
	}

	return user, accessToken, refreshToken, nil
}
