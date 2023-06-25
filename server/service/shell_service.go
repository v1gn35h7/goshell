package service

import (
	"github.com/google/uuid"
	respository "github.com/v1gn35h7/goshell/internal/repository"
	"github.com/v1gn35h7/goshell/pkg/goshell"
	"github.com/v1gn35h7/goshell/pkg/kafka"
	"github.com/v1gn35h7/goshell/pkg/logging"
)

type shellService interface {
	ExecuteCmd(cmd string) (string, error)
	ConnectToRemoteHost(hostId string) (bool, error)
	GetScripts(asset goshell.Asset) ([]*goshell.Script, error)
	SaveScripts(scriptPayload goshell.Script) (bool, error)
	//EndpointHeartBeat(hostId string) ([]execu)
	SendFragment(payload goshell.Fragment) (int32, error)
	SearchResults(query string) ([]*goshell.Output, error)
}

func (srvc service) ExecuteCmd(cmd string) (string, error) {
	return "ellow!!", nil
}

func (srvc service) ConnectToRemoteHost(hostId string) (bool, error) {
	return true, nil
}

func (srvc service) GetScripts(asset goshell.Asset) ([]*goshell.Script, error) {
	respository.AssetsRepository(logging.Logger()).UpdateAsset(asset)
	return respository.ScriptsRepository(logging.Logger()).GetScripts(asset.Agentid)
}

func (srvc service) SaveScripts(scriptPayload goshell.Script) (bool, error) {
	scriptPayload.Id = uuid.NewString()
	return respository.ScriptsRepository(logging.Logger()).AddScripts(scriptPayload)
}

func (srvc service) SendFragment(fragment goshell.Fragment) (int32, error) {
	//respository.AssetsRepository(logging.Logger()).UpdateAsset(asset)
	kafka.ProduceRecord(srvc, fragment)

	return int32(1), nil
}

func (srvc service) SearchResults(query string) ([]*goshell.Output, error) {
	return respository.ResultsRepository(logging.Logger()).SearchResults(query)
}
