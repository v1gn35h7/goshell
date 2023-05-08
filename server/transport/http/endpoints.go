package http

import (
	"github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	"github.com/goShell/server/service"
)

type shellEndpoints struct {
	executeCmdEndpoint    endpoint.Endpoint
	connectToHostEndpoint endpoint.Endpoint
	getAssetsEndpoint     endpoint.Endpoint
	getUsersEndpoint      endpoint.Endpoint
	addUserEndpoint       endpoint.Endpoint
}

func MakeEndpoints(srvc service.Service, logger log.Logger) shellEndpoints {
	endpoints := shellEndpoints{
		// Shell
		executeCmdEndpoint: makeExecuteCmdEndpointMiddleware(srvc, logger),
		//executeCmdEndpointMiddleware := logging.LoggingMiddleware(logger)(executeCmdEndpoint)

		connectToHostEndpoint: makeConnectToHostEndpointMiddleware(srvc, logger),
		//connectToHostEndpointMiddleware := logging.LoggingMiddleware(logger)(connectToHostEndpoint)

		//Assets
		getAssetsEndpoint: MakeGetAssetsEndpoint(srvc),
		//getAssetsEndpointMiddleware := logging.LoggingMiddleware(logger)(getAssetsEndpoint)

		//Users
		getUsersEndpoint: MakeGetUsersEndpoint(srvc),
		//getUsersEndpointMiddleware := logging.LoggingMiddleware(logger)(getUsersEndpoint)

		addUserEndpoint: MakeAddUsersEndpoint(srvc),
		//addUserEndpointMiddleware := logging.LoggingMiddleware(logger)(addUserEndpoint)
	}

	return endpoints
}
