package service

import (
	"fmt"
	"time"
)

func (middelware instrumentationServiceMiddleware) getEventsProto() (output bool, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "getEventsProto", "error", fmt.Sprint(err != nil)}
		middelware.requestCount.With(lvs...).Add(1)
		middelware.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return true, nil
}

func (middelware instrumentationServiceMiddleware) pushEvents() (output bool, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "pushEvents", "error", fmt.Sprint(err != nil)}
		middelware.requestCount.With(lvs...).Add(1)
		middelware.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return true, nil
}
