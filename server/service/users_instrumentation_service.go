package service

import (
	"fmt"
	"time"

	gomodels "github.com/v1gn35h7/goshell/pkg/goshell"
	"github.com/v1gn35h7/goshell/server/goshell"
)

func (middelware instrumentationServiceMiddleware) GetUsers() (output []*gomodels.Asset, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetAssets", "error", fmt.Sprint(err != nil)}
		middelware.requestCount.With(lvs...).Add(1)
		middelware.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = middelware.next.GetAssets()
	return
}

func (middelware instrumentationServiceMiddleware) AddUser(user goshell.User) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "AddUser", "error", fmt.Sprint(err != nil)}
		middelware.requestCount.With(lvs...).Add(1)
		middelware.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = middelware.next.AddUser(user)
	return
}
