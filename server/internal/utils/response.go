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
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func SendSuccessResponse(ctx *gin.Context, message string, data interface{}) {
	response := SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}

	ctx.JSON(http.StatusOK, gin.H{"response": response})
}

func SendErrorResponse(ctx *gin.Context, errorData ErrorResponse) {
	ctx.AbortWithStatusJSON(errorData.Code, gin.H{"response": errorData})
}
