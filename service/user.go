package service

import (
	"crud-app/domain"
	"crud-app/repository"
	"fmt"
)

func InsertUserData(user *domain.User) error {
	err := repository.NewUserRepository().InsertUser(user)

	if err != nil {
		fmt.Errorf("Error: %s", err)
		return err
	}

	return nil
}

func GetUser(userId int) (*domain.User, error) {
	user := &domain.User{}

	user, err := repository.NewUserRepository().GetUser(userId)
	if err != nil {
		return &domain.User{}, err
	}
	return user, nil
}

func GetAllUser() ([]domain.User, error) {
	user := []domain.User{}

	user, err := repository.NewUserRepository().GetAllUser()
	if err != nil {
		return []domain.User{}, err
	}
	return user, nil
}
