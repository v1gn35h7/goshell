package service

import (
	"fmt"
	"time"

	gomodels "github.com/v1gn35h7/goshell/pkg/goshell"
	"github.com/v1gn35h7/goshell/server/goshell"
)

func (m instrumentationMiddleware) GetUsers() (output []*gomodels.Asset, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetAssets", "error", fmt.Sprint(err != nil)}
		m.requestCount.With(lvs...).Add(1)
		m.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = m.next.GetAssets()
	return
}

func (m instrumentationMiddleware) AddUser(user goshell.User) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "AddUser", "error", fmt.Sprint(err != nil)}
		m.requestCount.With(lvs...).Add(1)
		m.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = m.next.AddUser(user)
	return
}
