package service

import (
	"time"

	"github.com/v1gn35h7/goshell/pkg/goshell"
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

func (middelware LoggingServiceMiddleware) GetScripts(asset goshell.Asset) ([]*goshell.Script, error) {
	defer func(tm time.Time) {
		middelware.logger.Log("Method", "GetScripts",
			"Time Since", time.Since(tm))
	}(time.Now())

	return middelware.next.GetScripts(asset)
}

func (middelware LoggingServiceMiddleware) SaveScripts(script goshell.Script) (bool, error) {
	defer func(tm time.Time) {
		middelware.logger.Log("Method", "SaveScripts",
			"Time Since", time.Since(tm))
	}(time.Now())

	return middelware.next.SaveScripts(script)
}
