package service

import (
	"fmt"
	"time"

	"github.com/v1gn35h7/goshell/pkg/goshell"
)

func (m instrumentationMiddleware) ExecuteCmd(cmd string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetAssets", "error", fmt.Sprint(err != nil)}
		m.requestCount.With(lvs...).Add(1)
		m.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = m.next.ExecuteCmd(cmd)
	return
}

func (m instrumentationMiddleware) ConnectToRemoteHost(hostId string) (output bool, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetAssets", "error", fmt.Sprint(err != nil)}
		m.requestCount.With(lvs...).Add(1)
		m.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = m.next.ConnectToRemoteHost(hostId)
	return
}

func (m instrumentationMiddleware) GetScripts(asset goshell.Asset) (output []*goshell.Script, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetScripts", "error", fmt.Sprint(err != nil)}
		m.requestCount.With(lvs...).Add(1)
		m.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return m.next.GetScripts(asset)
}

func (m instrumentationMiddleware) SaveScripts(script goshell.Script) (output bool, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "SaveScripts", "error", fmt.Sprint(err != nil)}
		m.requestCount.With(lvs...).Add(1)
		m.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = m.next.SaveScripts(script)
	return
}

func (m instrumentationMiddleware) SendFragment(payload goshell.Fragment) (ackw int32, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "SendFragment", "error", fmt.Sprint(err != nil)}
		m.requestCount.With(lvs...).Add(1)
		m.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return m.next.SendFragment(payload)
}

func (m instrumentationMiddleware) SearchResults(query string) (list []*goshell.Output, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "SearchResults", "error", fmt.Sprint(err != nil)}
		m.requestCount.With(lvs...).Add(1)
		m.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return m.next.SearchResults(query)
}
