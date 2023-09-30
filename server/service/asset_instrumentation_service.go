package service

import (
	"fmt"
	"time"

	"github.com/v1gn35h7/goshell/pkg/goshell"
)

func (m instrumentationMiddleware) GetAssets() (output []*goshell.Asset, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetAssets", "error", fmt.Sprint(err != nil)}
		m.requestCount.With(lvs...).Add(1)
		m.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = m.next.GetAssets()
	return
}
