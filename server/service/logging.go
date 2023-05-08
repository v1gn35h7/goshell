package service

import "github.com/go-kit/log"

type LoggingServiceMiddleware struct {
	logger log.Logger
	next   Service
}

func NewLoggingServiceMiddleware(logger log.Logger, next Service) LoggingServiceMiddleware {
	return LoggingServiceMiddleware{logger, next}
}
