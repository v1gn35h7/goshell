package service

import (
	"fmt"
	"time"
)

func (m instrumentationMiddleware) getEventsProto() (output bool, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "getEventsProto", "error", fmt.Sprint(err != nil)}
		m.requestCount.With(lvs...).Add(1)
		m.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return true, nil
}

func (m instrumentationMiddleware) pushEvents() (output bool, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "pushEvents", "error", fmt.Sprint(err != nil)}
		m.requestCount.With(lvs...).Add(1)
		m.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return true, nil
}
