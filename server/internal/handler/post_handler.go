package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sunilkkhadka/Forum/internal/service"
)

type PostHandler struct {
	PostService service.PostServiceI
}

func NewPostHandler(postService service.PostServiceI) *PostHandler {
	return &PostHandler{
		PostService: postService,
	}
}

func (handler *PostHandler) GetAllPosts(ctx *gin.Context) {
	posts, _ := handler.PostService.GetAllPosts()

	ctx.JSON(200, posts)
}
