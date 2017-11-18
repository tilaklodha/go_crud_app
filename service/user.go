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
