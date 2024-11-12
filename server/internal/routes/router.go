package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	api := router.Group("/api")
	RegisterV1Routes(api)
}
