package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/v1gn35h7/goshell/pkg/goshell"
)

// Shell serviec contracts
type executeCmdRequest struct {
	Command string `json:"command"`
}

type executeCmdResponse struct {
	Response string `json:"response"`
}

type connectToEndpointRequest struct {
	AgentId         string `json:"AgentId,omitempty"`
	HostName        string `json:"HostName,omitempty"`
	Platform        string `json:"Platform,omitempty"`
	OperatingSystem string `json:"OperatingSystem,omitempty"`
	Architecture    string `json:"Architecture,omitempty"`
}

type searchResultsRequest struct {
	Query string `json:"query"`
}

// ****************************************************************************************************
// Response structs
// ****************************************************************************************************
type connectToEndpointResponse struct {
	Error  string `json:"error"`
	Status string `json:"status"`
}

type saveScriptEndpointRequest struct {
	Title     string
	Script    string
	Platform  string `json:"platform"`
	Type      string
	Frequency string `json:"frequency"`
}

type saveScriptEndpointResponse struct {
	Error  string `json:"error"`
	Status string `json:"status"`
}

type getScriptsEndpointResponse struct {
	Error  string           `json:"error"`
	Status string           `json:"status"`
	List   []goshell.Script `json:"list"`
}

type searchResultsResponse struct {
	Error  string           `json:"error"`
	Status string           `json:"status"`
	List   []goshell.Output `json:"list"`
	Count  int              `json:"count"`
}

// ********************************************************************************************************************
// Scaffloding endpoints to transport
// *******************************************************************************************************************
func makeExecuteCmdTransport(endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(
		endpoint,
		decodeExecuteCmdRequest,
		encodeResponse,
		httpSrvOptions...,
	)
}

func makeConnectToHostEndpointTransport(endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(
		endpoint,
		decodeConnectToEndpointRequest,
		encodeResponse,
		httpSrvOptions...,
	)
}

func makeSaveScriptEndpointTransport(endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(
		endpoint,
		decodeSaveScriptEndpointRequest,
		encodeResponse,
		httpSrvOptions...,
	)
}

func makeGetScriptEndpointTransport(endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(
		endpoint,
		decodeConnectToEndpointRequest,
		encodeResponse,
		httpSrvOptions...,
	)
}

func makeSearchResultsEndpointTransport(endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(
		endpoint,
		decodeSearchResultsRequest,
		encodeResponse,
		httpSrvOptions...,
	)
}

// **************************************************************************************************************************
// Request utilities
// **************************************************************************************************************************
func decodeExecuteCmdRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	var req executeCmdRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func decodeConnectToEndpointRequest(ctx context.Context, request *http.Request) (interface{}, error) {

	req := connectToEndpointRequest{
		AgentId: request.URL.Query().Get("hostId"),
	}

	return req, nil
}

func decodeSaveScriptEndpointRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	var req saveScriptEndpointRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func decodeSearchResultsRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	req := searchResultsRequest{
		Query: request.URL.Query().Get("query"),
	}

	return req, nil
}
