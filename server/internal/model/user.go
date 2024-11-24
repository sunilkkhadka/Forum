package model

import (
	"database/sql"
	"time"
)

type BasicUser struct {
	Email    string
	Password string
}

type User struct {
	ID             int
	FirstName      sql.NullString
	LastName       sql.NullString
	Email          string
	PasswordHash   string
	ProfilePicture sql.NullString
	CreatedAt      time.Time
	UpdatedAt      sql.NullTime
}
