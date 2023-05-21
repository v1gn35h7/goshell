package service

import (
	respository "github.com/v1gn35h7/goshell/internal/repository"
	"github.com/v1gn35h7/goshell/pkg/goshell"
	"github.com/v1gn35h7/goshell/pkg/logging"
)

type assetService interface {
	GetAssets() ([]*goshell.Asset, error)
}

func (srvc service) GetAssets() ([]*goshell.Asset, error) {
	return respository.AssetsRepository(logging.Logger()).GetAssets("")
}
