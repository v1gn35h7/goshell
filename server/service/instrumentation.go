package service

import "github.com/go-kit/kit/metrics"

type instrumentationServiceMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           Service
}

func NewInstrumentationServiceMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram, next Service) instrumentationServiceMiddleware {
	return instrumentationServiceMiddleware{requestCount, requestLatency, next}
}
