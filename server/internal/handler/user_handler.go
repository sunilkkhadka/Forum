package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunilkkhadka/Forum/internal/dto"
	"github.com/sunilkkhadka/Forum/internal/service"
	"github.com/sunilkkhadka/Forum/internal/utils"
	"github.com/sunilkkhadka/Forum/internal/utils/auth"
)

type UserHandler struct {
	UserService service.UserServiceI
}

func NewUserHandler(userService service.UserServiceI) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (handler *UserHandler) RegisterUserHandler(ctx *gin.Context) {
	var registerRequest *dto.BasicUserRequest

	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
		utils.SendErrorResponse(ctx, utils.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Credentials",
		})
		return
	}

	if err := dto.ValidateBasicUser(registerRequest); err != nil {
		utils.SendErrorResponse(ctx, utils.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err := handler.UserService.RegisterUser(registerRequest)
	if err != nil {
		var userErr *utils.UserFacingError
		if errors.As(err, &userErr) {
			utils.SendErrorResponse(ctx, utils.ErrorResponse(*userErr))
		} else {
			log.Printf("Server error occurred: %v", err)
			utils.SendErrorResponse(ctx, utils.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "An unexpected error occurred",
			})
		}
		return
	}

	utils.SendSuccessResponse(ctx, "Registration Successful", "")
}

func (handler *UserHandler) LoginUserHandler(ctx *gin.Context) {
	var loginRequest *dto.BasicUserRequest

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		utils.SendErrorResponse(ctx, utils.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Credentials",
		})
		return
	}

	if err := dto.ValidateBasicUser(loginRequest); err != nil {
		utils.SendErrorResponse(ctx, utils.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	user, accessToken, refreshToken, err := handler.UserService.LoginUser(loginRequest)

	if err != nil {
		var userErr *utils.UserFacingError
		if errors.As(err, &userErr) {
			utils.SendErrorResponse(ctx, utils.ErrorResponse(*userErr))
		} else {
			log.Printf("Server error occurred: %v", err)
			utils.SendErrorResponse(ctx, utils.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "An unexpected error occurred",
			})
		}
		return
	}

	ctx.SetCookie("access_token", accessToken, int(auth.JwtConf.JwtAccessTokenExpirationTime.Seconds()), "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refreshToken, int(auth.JwtConf.JwtRefreshTokenExpirationTime.Seconds()), "/api/v1/auth", "localhost", false, true)

	utils.SendSuccessResponse(ctx, "Logged In Successfully", map[string]any{
		"email": user.Email,
	})

}
