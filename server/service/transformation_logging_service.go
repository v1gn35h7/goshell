package service

import "time"

func (middelware LoggingServiceMiddleware) getEventsProto() (bool, error) {
	defer func(tm time.Time) {
		middelware.logger.Log("Method", "getEventsProto",
			"Time Since", time.Since(tm))
	}(time.Now())

	return true, nil
}

func (middelware LoggingServiceMiddleware) pushEvents() (bool, error) {
	defer func(tm time.Time) {
		middelware.logger.Log("Method", "pushEvents",
			"Time Since", time.Since(tm))
	}(time.Now())

	return true, nil
}
