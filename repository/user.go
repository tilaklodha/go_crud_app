package repository

import (
	"crud-app/appcontext"
	"crud-app/domain"
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	insertQuery   = "INSERT INTO users (first_name, last_name, city, created_at) VALUES ($1, $2, $3, $4)"
	fetchQuery    = "Select first_name, last_name, city from users where id = $1"
	fetchAllQuery = "Select first_name, last_name, city from users"
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

func (ur *userRepository) GetUser(userId int) (*domain.User, error) {
	data := ur.db.QueryRow(fmt.Sprintf(fetchQuery), userId)
	user := domain.User{}
	err := data.Scan(&user.FirstName, &user.LastName, &user.City)
	if err != nil {
		return &domain.User{}, err
	}
	return &user, nil
}

func (ur *userRepository) GetAllUser() ([]domain.User, error) {
	users := []domain.User{}
	data, err := ur.db.Query(fmt.Sprintf(fetchAllQuery))
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	for data.Next() {
		user := domain.User{}
		err := data.Scan(&user.FirstName, &user.LastName, &user.City)
		if err != nil {
			user = domain.User{}
		}
		users = append(users, user)
	}
	return users, nil
}
