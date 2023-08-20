package grpc

import (
	"context"

	log "github.com/go-kit/kit/log"

	"github.com/go-kit/kit/endpoint"
	"github.com/v1gn35h7/goshell/pkg/goshell"
	"github.com/v1gn35h7/goshell/server/logging"
	"github.com/v1gn35h7/goshell/server/pb"
	"github.com/v1gn35h7/goshell/server/service"
)

// Endpoints creators
func makeExecutGetScriptsEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(*pb.ShellRequest)
		asset := goshell.Asset{Agentid: string(req.AgentId),
			Hostname:        string(req.HostName),
			Platform:        string(req.Platform),
			Operatingsystem: string(req.OperatingSystem),
			Architecture:    string(req.Architecture),
		}

		res, err := srvc.GetScripts(asset)

		return GetScriptResponse{Scripts: res}, err
	}
}

func MakeGetScriptsEndpointMiddleware(srvc service.Service, logger log.Logger) endpoint.Endpoint {
	getScriptEndpoint := makeExecutGetScriptsEndpoint(srvc)
	return logging.LoggingMiddleware(logger)(getScriptEndpoint)
}

func makeSendFragmentEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(*pb.ShellFragmentRquest)
		payload := goshell.Fragment{
			Outputs: make([]goshell.Output, 0),
		}

		for _, op := range req.Outputs {
			otp := goshell.Output{
				Agentid:  op.AgentId,
				Hostname: op.HostName,
				Scriptid: op.ScriptId,
				Output:   op.Output,
			}
			payload.Outputs = append(payload.Outputs, otp)
		}

		res, err := srvc.SendFragment(payload)

		return FragmentResponse{Awknowledgement: res}, err
	}
}

func MakeSendFragmentsEndpointMiddleware(srvc service.Service, logger log.Logger) endpoint.Endpoint {
	sendFragmentEndpoint := makeSendFragmentEndpoint(srvc)
	return logging.LoggingMiddleware(logger)(sendFragmentEndpoint)
}
