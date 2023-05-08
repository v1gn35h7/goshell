package service

type Service interface {
	assetService
	userService
	shellService
}

type service struct{}

func New() service {
	return service{}
}
