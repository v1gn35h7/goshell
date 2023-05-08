package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// Shell serviec contracts
type executeCmdRequest struct {
	Command string `json:"command"`
}

type executeCmdResponse struct {
	Response string `json:"response"`
}

type connectToEndpointRequest struct {
	hostId string
}

type connectToEndpointResponse struct {
	Error  string `json:"error"`
	Status string `json:"status"`
}

// Scaffloding endpoints to transport
func makeExecuteCmdTransport(endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(
		endpoint,
		decodeExecuteCmdRequest,
		encodeExecuteCmdRequest,
		httpSrvOptions...,
	)
}

func makeConnectToHostEndpointTransport(endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(
		endpoint,
		decodeConnectToEndpointRequest,
		encodeConnectToEndpointResponse,
		httpSrvOptions...,
	)
}

// Request utilities
func decodeExecuteCmdRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	var req executeCmdRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func encodeExecuteCmdRequest(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeConnectToEndpointRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	var req connectToEndpointRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func encodeConnectToEndpointResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
