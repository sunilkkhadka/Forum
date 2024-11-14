package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	v1 "github.com/sunilkkhadka/Forum/internal/api/v1"
	"github.com/sunilkkhadka/Forum/internal/handler"
	"github.com/sunilkkhadka/Forum/internal/repository"
	"github.com/sunilkkhadka/Forum/internal/service"
)

func RegisterV1Routes(rg *gin.RouterGroup, db *sql.DB) {
	route := rg.Group("/v1")

	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)

	v1.PostRoutes(route, postHandler)
}
