package service

import (
	"fmt"
	"time"

	"github.com/v1gn35h7/goshell/pkg/goshell"
)

func (middelware instrumentationServiceMiddleware) ExecuteCmd(cmd string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetAssets", "error", fmt.Sprint(err != nil)}
		middelware.requestCount.With(lvs...).Add(1)
		middelware.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = middelware.next.ExecuteCmd(cmd)
	return
}

func (middelware instrumentationServiceMiddleware) ConnectToRemoteHost(hostId string) (output bool, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetAssets", "error", fmt.Sprint(err != nil)}
		middelware.requestCount.With(lvs...).Add(1)
		middelware.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = middelware.next.ConnectToRemoteHost(hostId)
	return
}

func (middelware instrumentationServiceMiddleware) GetScripts(asset goshell.Asset) (output []*goshell.Script, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetScripts", "error", fmt.Sprint(err != nil)}
		middelware.requestCount.With(lvs...).Add(1)
		middelware.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return middelware.next.GetScripts(asset)
}

func (middelware instrumentationServiceMiddleware) SaveScripts(script goshell.Script) (output bool, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "SaveScripts", "error", fmt.Sprint(err != nil)}
		middelware.requestCount.With(lvs...).Add(1)
		middelware.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = middelware.next.SaveScripts(script)
	return
}

func (middelware instrumentationServiceMiddleware) SendFragment(payload goshell.Fragment) (ackw int32, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "SendFragment", "error", fmt.Sprint(err != nil)}
		middelware.requestCount.With(lvs...).Add(1)
		middelware.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return middelware.next.SendFragment(payload)
}

func (middelware instrumentationServiceMiddleware) SearchResults(query string) (list []*goshell.Output, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "SearchResults", "error", fmt.Sprint(err != nil)}
		middelware.requestCount.With(lvs...).Add(1)
		middelware.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return middelware.next.SearchResults(query)
}
