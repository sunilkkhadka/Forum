package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/sunilkkhadka/Forum/internal/model"
	"github.com/sunilkkhadka/Forum/internal/utils"
)

type UserRepositoryI interface {
	LoginUser(user *model.BasicUser) error
	RegisterUser(user *model.BasicUser) error
	GetUserByEmail(email string) (*model.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryI {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) RegisterUser(user *model.BasicUser) error {

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

func (repo *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	stmt, err := repo.db.Prepare("SELECT id, first_name, last_name, email, password_hash, profile_picture, created_at, updated_at FROM users WHERE email = ? AND deleted_at IS NULL")
	if err != nil {
		return nil, fmt.Errorf("cannot prepare statement to find user by email: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.PasswordHash, &user.ProfilePicture, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utils.ErrUserDoNotExist
		}
		return nil, fmt.Errorf("couldn't query users: %w", err)
	}

	return &user, nil
}

func (repo *UserRepository) LoginUser(user *model.BasicUser) error {
	return nil
}
