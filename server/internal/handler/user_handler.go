package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunilkkhadka/Forum/internal/dto"
	"github.com/sunilkkhadka/Forum/internal/service"
	"github.com/sunilkkhadka/Forum/internal/utils"
)

type UserHandler struct {
	UserService service.UserServiceI
}

func NewUserHandler(userService service.UserServiceI) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (handler *UserHandler) RegisterUser(ctx *gin.Context) {
	var registerUser *dto.RegisterUserRequest

	if err := ctx.ShouldBindJSON(&registerUser); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := dto.ValidateRegisterUser(registerUser); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	utils.SendSuccessResponse(ctx, "Success bhayo hai", registerUser)
}
