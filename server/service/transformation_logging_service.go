package service

import "time"

func (m LoggingMiddleware) getEventsProto() (bool, error) {
	defer func(tm time.Time) {
		m.logger.Log("Method", "getEventsProto",
			"Time Since", time.Since(tm))
	}(time.Now())

	return true, nil
}

func (m LoggingMiddleware) pushEvents() (bool, error) {
	defer func(tm time.Time) {
		m.logger.Log("Method", "pushEvents",
			"Time Since", time.Since(tm))
	}(time.Now())

	return true, nil
}
