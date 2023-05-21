package service

import (
	"time"

	"github.com/v1gn35h7/goshell/pkg/goshell"
)

func (middelware LoggingServiceMiddleware) GetAssets() ([]*goshell.Asset, error) {
	defer func(tm time.Time) {
		middelware.logger.Log("Method", "GetAssets",
			"Time Since", time.Since(tm))
	}(time.Now())

	return middelware.next.GetAssets()
}
