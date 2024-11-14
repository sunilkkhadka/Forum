package repository

import (
	"database/sql"

	"github.com/sunilkkhadka/Forum/internal/model"
)

type PostRepositoryI interface {
	GetAllPosts() (*[]model.Post, error)
}

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepositoryI {
	return &PostRepository{
		db: db,
	}
}

func (repo *PostRepository) GetAllPosts() (*[]model.Post, error) {
	posts := &[]model.Post{
		{
			ID:        1,
			Content:   "Some Content Here",
			Title:     "React Basics",
			UserID:    1,
			Tags:      []string{"web", "react"},
			CreatedAt: "Nov 14, 2024 - Thursday",
			UpdatedAt: "",
			DeletedAt: "",
		},
		{
			ID:        2,
			Content:   "Some Other Content Here",
			Title:     "Next Js",
			UserID:    1,
			Tags:      []string{"web", "react"},
			CreatedAt: "Nov 14, 2024 - Thursday",
			UpdatedAt: "",
			DeletedAt: "",
		},
	}

	return posts, nil

}
