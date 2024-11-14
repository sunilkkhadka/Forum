package service

import (
	"github.com/sunilkkhadka/Forum/internal/model"
	"github.com/sunilkkhadka/Forum/internal/repository"
)

type PostServiceI interface {
	GetAllPosts() (*[]model.Post, error)
}

type PostService struct {
	PostRepo repository.PostRepositoryI
}

func NewPostService(postRepo repository.PostRepositoryI) PostServiceI {
	return &PostService{
		PostRepo: postRepo,
	}
}

func (service *PostService) GetAllPosts() (*[]model.Post, error) {
	posts, _ := service.PostRepo.GetAllPosts()
	return posts, nil
}
