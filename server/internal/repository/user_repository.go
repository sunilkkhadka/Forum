package repository

import (
	"database/sql"
	"fmt"

	"github.com/sunilkkhadka/Forum/internal/model"
)

type UserRepositoryI interface {
	RegisterUser(user *model.RegisterUser) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryI {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) RegisterUser(user *model.RegisterUser) error {

	stmt, err := repo.db.Prepare("INSERT INTO users (email, password_hash) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("cannot prepare statement to register user: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("cannot create user: %w", err)
	}

	_, err = result.LastInsertId()
	if err != nil {
		return fmt.Errorf("wtf is going on")
	}

	return nil
}
