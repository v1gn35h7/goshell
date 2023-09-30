package http

import (
	"github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	"github.com/v1gn35h7/goshell/server/service"
)

type goshellEndpoints struct {
	executeCmdEndpoint    endpoint.Endpoint
	connectToHostEndpoint endpoint.Endpoint
	getAssetsEndpoint     endpoint.Endpoint
	getUsersEndpoint      endpoint.Endpoint
	addUserEndpoint       endpoint.Endpoint
	saveScriptEndpoint    endpoint.Endpoint
	getScriptsEndpoint    endpoint.Endpoint
	searchResultsEndpoint endpoint.Endpoint
}

func MakeEndpoints(srvc service.Service, logger log.Logger) goshellEndpoints {
	endpoints := goshellEndpoints{
		// Shell
		executeCmdEndpoint: makeExecuteCmdEndpointMiddleware(srvc, logger),

		saveScriptEndpoint: makeSaveScriptsEndpointMiddleware(srvc, logger),

		searchResultsEndpoint: makeSearchResultsEndpointMiddleware(srvc, logger),
		//executeCmdEndpointMiddleware := logging.LoggingMiddleware(logger)(executeCmdEndpoint)

		getScriptsEndpoint: makeGetScriptsEndpointMiddleware(srvc, logger),

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
