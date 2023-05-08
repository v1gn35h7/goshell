package service

import (
	"time"
)

func (middelware LoggingServiceMiddleware) GetAssets() ([]string, error) {
	defer func(tm time.Time) {
		middelware.logger.Log("Method", "GetAssets",
			"Time Since", time.Since(tm))
	}(time.Now())

	return middelware.next.GetAssets()
}
