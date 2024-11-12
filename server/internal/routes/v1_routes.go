package routes

import (
	"github.com/gin-gonic/gin"
	ver1 "github.com/sunilkkhadka/Forum/internal/handlers/v1"
)

func RegisterV1Routes(rg *gin.RouterGroup) {
	v1 := rg.Group("/v1")
	v1.GET("/healthcheck", ver1.HealthCheck)
}
