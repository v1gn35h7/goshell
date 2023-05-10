package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/v1gn35h7/goshell/server/service"
)

var (
	httpSrvOptions []httptransport.ServerOption
)

func MakeHandlers(srvc service.Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := MakeEndpoints(srvc, logger)
	httpSrvOptions = []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}
	r.PathPrefix("/app").Handler(MakeFrontEndHandler())
	r.Handle("/exec", makeExecuteCmdTransport(e.executeCmdEndpoint)).Methods("GET")
	r.Handle("/connect", makeConnectToHostEndpointTransport(e.connectToHostEndpoint)).Methods("GET")
	r.Handle("/metrics", promhttp.Handler()).Methods("GET")
	r.Handle("/assets", getAssetsEnpointTransport(e.getAssetsEndpoint)).Methods("GET")
	r.Handle("/users", getUsersEndpointTransport(e.getUsersEndpoint)).Methods("GET").Name("get_uses")
	r.Handle("/users/add", addUserEndpointTransport(e.addUserEndpoint)).Methods("POST").Name("add_user")
	r.PathPrefix("/").Handler(MakeFrontEndHandler())
	return r
}

// Common
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	// switch err {
	// case ErrNotFound:
	// 	return http.StatusNotFound
	// case ErrAlreadyExists, ErrInconsistentIDs:
	// 	return http.StatusBadRequest
	// default:
	// 	return http.StatusInternalServerError
	// }
	return http.StatusInternalServerError
}
