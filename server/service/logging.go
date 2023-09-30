package service

import "github.com/go-kit/log"

type LoggingMiddleware struct {
	logger log.Logger
	next   Service
}

func NewLoggingMiddleware(logger log.Logger, next Service) LoggingMiddleware {
	return LoggingMiddleware{logger, next}
}
