package service

import (
	"time"

	"github.com/v1gn35h7/goshell/pkg/goshell"
)

func (m LoggingMiddleware) ExecuteCmd(cmd string) (string, error) {
	defer func(tm time.Time) {
		m.logger.Log("Method", "ExecuteCmd",
			"Time Since", time.Since(tm))
	}(time.Now())

	return m.next.ExecuteCmd(cmd)
}

func (m LoggingMiddleware) ConnectToRemoteHost(hostId string) (bool, error) {
	defer func(tm time.Time) {
		m.logger.Log("Method", "ConnectToRemoteHost",
			"Time Since", time.Since(tm))
	}(time.Now())

	return m.next.ConnectToRemoteHost(hostId)
}

func (m LoggingMiddleware) GetScripts(asset goshell.Asset) ([]*goshell.Script, error) {
	defer func(tm time.Time) {
		m.logger.Log("Method", "GetScripts",
			"Time Since", time.Since(tm))
	}(time.Now())

	return m.next.GetScripts(asset)
}

func (m LoggingMiddleware) SaveScripts(script goshell.Script) (bool, error) {
	defer func(tm time.Time) {
		m.logger.Log("Method", "SaveScripts",
			"Time Since", time.Since(tm))
	}(time.Now())

	return m.next.SaveScripts(script)
}

func (m LoggingMiddleware) SendFragment(payload goshell.Fragment) (int32, error) {
	defer func(tm time.Time) {
		m.logger.Log("Method", "SendFragment",
			"Time Since", time.Since(tm))
	}(time.Now())

	return m.next.SendFragment(payload)
}

func (m LoggingMiddleware) SearchResults(query string) ([]*goshell.Output, error) {
	defer func(tm time.Time) {
		m.logger.Log("Method", "Search Results",
			"Time Since", time.Since(tm))
	}(time.Now())

	return m.next.SearchResults(query)
}
