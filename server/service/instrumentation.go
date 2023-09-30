package service

import "github.com/go-kit/kit/metrics"

type instrumentationMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           Service
}

func NewInstrumentationMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram, next Service) instrumentationMiddleware {
	return instrumentationMiddleware{requestCount, requestLatency, next}
}
