package service

import (
	"time"

	gomodel "github.com/v1gn35h7/goshell/pkg/goshell"
	"github.com/v1gn35h7/goshell/server/goshell"
)

func (middelware LoggingServiceMiddleware) GetUsers() ([]*gomodel.Asset, error) {
	defer func(tm time.Time) {
		middelware.logger.Log("Method", "GetUsers",
			"Time Since", time.Since(tm))
	}(time.Now())

	return middelware.next.GetUsers()
}

func (middelware LoggingServiceMiddleware) AddUser(user goshell.User) (string, error) {
	defer func(tm time.Time) {
		middelware.logger.Log("Method", "AddUser",
			"Time Since", time.Since(tm))
	}(time.Now())

	return middelware.next.AddUser(user)
}
