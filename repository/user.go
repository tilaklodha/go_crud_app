package repository

import (
	"crud-app/appcontext"
	"crud-app/domain"
	"database/sql"
	"fmt"
	"time"
)

const (
	insertQuery = "INSERT INTO users (first_name, last_name, city, created_at) VALUES ($1, $2, $3, $4)"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository() *userRepository {
	return &userRepository{
		db: appcontext.GetDB(),
	}
}

func (ur *userRepository) InsertUser(user *domain.User) error {
	_, err := ur.db.Exec(insertQuery, user.FirstName, user.LastName, user.City, time.Now())
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}

	return err
}
