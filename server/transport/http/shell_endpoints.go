package http

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	"github.com/v1gn35h7/goshell/pkg/goshell"
	"github.com/v1gn35h7/goshell/server/logging"
	"github.com/v1gn35h7/goshell/server/service"
)

// Endpoints creators
func makeExecuteCmdEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(executeCmdRequest)
		res, err := srvc.ExecuteCmd(req.Command)

		return executeCmdResponse{Response: res}, err
	}
}

func makeExecuteCmdEndpointMiddleware(srvc service.Service, logger log.Logger) endpoint.Endpoint {
	executeCmdEndpoint := makeExecuteCmdEndpoint(srvc)
	return logging.LoggingMiddleware(logger)(executeCmdEndpoint)
}

func makeConnectToHostEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(connectToEndpointRequest)
		res, err := srvc.ConnectToRemoteHost(req.AgentId)
		status := "SUCCESS"
		if !res {
			status = "FAILED"
			return connectToEndpointResponse{Error: "Failed to connect to the endpoint", Status: status}, err
		}

		return connectToEndpointResponse{Error: "", Status: status}, nil
	}
}

func makeConnectToHostEndpointMiddleware(srvc service.Service, logger log.Logger) endpoint.Endpoint {
	connectToHostEndpoint := makeConnectToHostEndpoint(srvc)
	return logging.LoggingMiddleware(logger)(connectToHostEndpoint)
}

func makeSaveScriptEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(saveScriptEndpointRequest)
		script := goshell.Script{
			Title:    req.Title,
			Platform: req.Platform,
			Type:     req.Type,
			Script:   req.Script,
		}
		res, err := srvc.SaveScripts(script)
		status := "SUCCESS"
		if !res {
			status = "FAILED"
			return saveScriptEndpointResponse{Error: "Failed to save script", Status: status}, err
		}

		return connectToEndpointResponse{Error: "", Status: status}, nil
	}
}

func makeSaveScriptsEndpointMiddleware(srvc service.Service, logger log.Logger) endpoint.Endpoint {
	saveScriptEndpoint := makeSaveScriptEndpoint(srvc)
	return logging.LoggingMiddleware(logger)(saveScriptEndpoint)
}

func makeGetScriptsEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(connectToEndpointRequest)
		asset := goshell.Asset{
			AgentId:         req.AgentId,
			Platform:        req.Platform,
			HostName:        req.HostName,
			Architecture:    req.Architecture,
			OperatingSystem: req.OperatingSystem,
		}

		list, err := srvc.GetScripts(asset)
		status := "SUCCESS"
		if err != nil {
			status = "FAILED"
			return getScriptsEndpointResponse{Error: "Failed to get scripts", Status: status}, err
		}

		// TODO: refactor this code
		data := make([]goshell.Script, 0)
		for _, v := range list {
			data = append(data, *v)
		}

		return getScriptsEndpointResponse{Error: "", Status: status, List: data}, nil
	}
}

func makeGetScriptsEndpointMiddleware(srvc service.Service, logger log.Logger) endpoint.Endpoint {
	getScriptEndpoint := makeGetScriptsEndpoint(srvc)
	return logging.LoggingMiddleware(logger)(getScriptEndpoint)
}
