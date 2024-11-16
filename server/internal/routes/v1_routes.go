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

	// Repositories
	postRepository := repository.NewPostRepository(db)
	userRepository := repository.NewUserRepository(db)

	// Services
	postService := service.NewPostService(postRepository)
	userService := service.NewUserService(userRepository)

	// Handlers
	postHandler := handler.NewPostHandler(postService)
	userHandler := handler.NewUserHandler(userService)

	// Routes
	v1.PostRoutes(route, postHandler)
	v1.UserRoutes(route, userHandler)
}
