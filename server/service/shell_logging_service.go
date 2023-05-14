package service

import (
	"time"

	"github.com/v1gn35h7/goshell/server/goshell"
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

func (middelware LoggingServiceMiddleware) GetScripts(agentId string) ([]*goshell.ShellScript, error) {
	defer func(tm time.Time) {
		middelware.logger.Log("Method", "GetScripts",
			"Time Since", time.Since(tm))
	}(time.Now())

	return middelware.next.GetScripts(agentId)
}
