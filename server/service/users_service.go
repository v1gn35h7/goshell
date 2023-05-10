package service

import (
	"fmt"

	"github.com/v1gn35h7/goshell/server/goshell"
)

type userService interface {
	GetUsers() ([]string, error)
	AddUser(goshell.User) (string, error)
}

func (srvc service) GetUsers() ([]string, error) {
	return []string{"Admin", "Sonu", "Sekar"}, nil
}

func (srvc service) AddUser(user goshell.User) (string, error) {
	return fmt.Sprintf("Added user %s", user.Username), nil
}
