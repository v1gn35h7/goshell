package service

import (
	"fmt"

	gomodel "github.com/v1gn35h7/goshell/pkg/goshell"
	"github.com/v1gn35h7/goshell/server/goshell"
)

type userService interface {
	GetUsers() ([]*gomodel.Asset, error)
	AddUser(goshell.User) (string, error)
}

func (srvc service) GetUsers() ([]*gomodel.Asset, error) {
	return make([]*gomodel.Asset, 0), nil
}

func (srvc service) AddUser(user goshell.User) (string, error) {
	return fmt.Sprintf("Added user %s", user.Username), nil
}
