package dto

import "github.com/go-playground/validator/v10"

type RegisterUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func ValidateRegisterUser(data *RegisterUserRequest) error {
	validate := validator.New()
	return validate.Struct(data)
}
