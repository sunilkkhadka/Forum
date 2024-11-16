package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sunilkkhadka/Forum/internal/handler"
)

func UserRoutes(rg *gin.RouterGroup, userHandler *handler.UserHandler) {
	auth_routes := rg.Group("/auth")

	auth_routes.POST("/register", userHandler.RegisterUser)
}
