package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sunilkkhadka/Forum/internal/handler"
	"github.com/sunilkkhadka/Forum/internal/middleware"
)

func PostRoutes(rg *gin.RouterGroup, postHandler *handler.PostHandler) {
	route := rg.Group("/post")

	route.Use(middleware.AuthMiddleware())
	route.GET("/all", postHandler.GetAllPosts)

}
