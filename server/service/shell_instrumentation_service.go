package service

import (
	"fmt"
	"time"

	"github.com/v1gn35h7/goshell/server/goshell"
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

func (middelware instrumentationServiceMiddleware) GetScripts(agentId string) (output []*goshell.ShellScript, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetScripts", "error", fmt.Sprint(err != nil)}
		middelware.requestCount.With(lvs...).Add(1)
		middelware.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = middelware.next.GetScripts(agentId)
	return
}
