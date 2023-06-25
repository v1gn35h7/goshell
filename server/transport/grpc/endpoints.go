package grpc

import (
	"github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	"github.com/v1gn35h7/goshell/server/service"
)

type grpcEndpoints struct {
	GetScriptEndpoint    endpoint.Endpoint
	SendFragmentEndpoint endpoint.Endpoint
}

func MakeGrpcEndpoints(srvc service.Service, logger log.Logger) grpcEndpoints {

	return grpcEndpoints{
		GetScriptEndpoint:    MakeGetScriptsEndpointMiddleware(srvc, logger),
		SendFragmentEndpoint: MakeSendFragmentsEndpointMiddleware(srvc, logger),
	}

}
