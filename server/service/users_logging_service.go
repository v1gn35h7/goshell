package service

import (
	"time"

	gomodel "github.com/v1gn35h7/goshell/pkg/goshell"
	"github.com/v1gn35h7/goshell/server/goshell"
)

func (m LoggingMiddleware) GetUsers() ([]*gomodel.Asset, error) {
	defer func(tm time.Time) {
		m.logger.Log("Method", "GetUsers",
			"Time Since", time.Since(tm))
	}(time.Now())

	return m.next.GetUsers()
}

func (m LoggingMiddleware) AddUser(user goshell.User) (string, error) {
	defer func(tm time.Time) {
		m.logger.Log("Method", "AddUser",
			"Time Since", time.Since(tm))
	}(time.Now())

	return m.next.AddUser(user)
}
