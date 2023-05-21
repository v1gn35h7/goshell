package service

import (
	"fmt"
	"time"

	"github.com/v1gn35h7/goshell/pkg/goshell"
)

func (middelware instrumentationServiceMiddleware) GetAssets() (output []*goshell.Asset, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetAssets", "error", fmt.Sprint(err != nil)}
		middelware.requestCount.With(lvs...).Add(1)
		middelware.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = middelware.next.GetAssets()
	return
}
