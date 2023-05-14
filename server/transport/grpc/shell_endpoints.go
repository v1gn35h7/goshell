package grpc

import (
	"context"

	log "github.com/go-kit/kit/log"

	"github.com/go-kit/kit/endpoint"
	"github.com/v1gn35h7/goshell/server/logging"
	"github.com/v1gn35h7/goshell/server/service"
)

// Endpoints creators
func makeExecutGetScriptsEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(GetScriptRequest)
		res, err := srvc.GetScripts(req.AgentId)

		return GetScriptResponse{Scripts: res}, err
	}
}

func MakeGetScriptsEndpointMiddleware(srvc service.Service, logger log.Logger) endpoint.Endpoint {
	getScriptEndpoint := makeExecutGetScriptsEndpoint(srvc)
	return logging.LoggingMiddleware(logger)(getScriptEndpoint)
}
