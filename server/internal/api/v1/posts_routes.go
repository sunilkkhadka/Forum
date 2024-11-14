package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sunilkkhadka/Forum/internal/handler"
)

func PostRoutes(rg *gin.RouterGroup, postHandler *handler.PostHandler) {
	route := rg.Group("/post")

	route.GET("/all", postHandler.GetAllPosts)

}
