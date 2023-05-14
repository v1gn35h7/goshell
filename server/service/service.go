package service

import "github.com/go-kit/log"

type Service interface {
	assetService
	userService
	shellService
	transformationService
}

type service struct {
	logger log.Logger
}

func New(logger log.Logger) service {
	return service{logger: logger}
}
