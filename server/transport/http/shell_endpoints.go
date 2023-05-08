package http

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	"github.com/goShell/server/logging"
	"github.com/goShell/server/service"
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
		res, err := srvc.ConnectToRemoteHost(req.hostId)
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
