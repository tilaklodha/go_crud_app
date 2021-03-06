package service

import (
	"fmt"
	"go_crud_app/domain"
	"go_crud_app/repository"
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

func DeleteUser(userId int) error {
	err := repository.NewUserRepository().DeleteUser(userId)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *domain.User, userId int) (*domain.User, error) {
	updateUser := &domain.User{}

	updateUser, err := repository.NewUserRepository().UpdateUser(user, userId)
	if err != nil {
		return &domain.User{}, err
	}
	return updateUser, nil
}
