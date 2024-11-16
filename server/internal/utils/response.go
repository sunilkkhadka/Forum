package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func SendSuccessResponse(ctx *gin.Context, message string, data interface{}) {
	response := SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}

	ctx.JSON(http.StatusOK, response)
}

func SendErrorResponse(ctx *gin.Context, code int, message string) {
	response := ErrorResponse{
		Status:  "error",
		Message: message,
	}
	ctx.AbortWithStatusJSON(code, response)
}
