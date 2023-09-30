package service

import (
	"time"

	"github.com/v1gn35h7/goshell/pkg/goshell"
)

func (m LoggingMiddleware) GetAssets() ([]*goshell.Asset, error) {
	defer func(tm time.Time) {
		m.logger.Log("Method", "GetAssets",
			"Time Since", time.Since(tm))
	}(time.Now())

	return m.next.GetAssets()
}
