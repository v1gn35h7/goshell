package service

import (
	"time"
)

func (middelware LoggingServiceMiddleware) ExecuteCmd(cmd string) (string, error) {
	defer func(tm time.Time) {
		middelware.logger.Log("Method", "ExecuteCmd",
			"Time Since", time.Since(tm))
	}(time.Now())

	return middelware.next.ExecuteCmd(cmd)
}

func (middelware LoggingServiceMiddleware) ConnectToRemoteHost(hostId string) (bool, error) {
	defer func(tm time.Time) {
		middelware.logger.Log("Method", "ConnectToRemoteHost",
			"Time Since", time.Since(tm))
	}(time.Now())

	return middelware.next.ConnectToRemoteHost(hostId)
}
