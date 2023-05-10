package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/v1gn35h7/goshell/server/goshell"
)

// Get Users
type getUsersRequest struct {
	Query string `json:"query"`
}

type getUsersResponse struct {
	List  []string `json:"list"`
	Count int64    `json:"count"`
}

// Add user
type addUserRequest struct {
	User goshell.User `json:"user"`
}

type addUserResponse struct {
	Status string `json:"status"`
}

// Endpoint Tranports
func addUserEndpointTransport(endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(
		endpoint,
		decodeAddUsersRequest,
		encodeResponse,
		httpSrvOptions...,
	)
}

func getUsersEndpointTransport(endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(
		endpoint,
		decodeGetAssetsRequest,
		encodeResponse,
		httpSrvOptions...,
	)
}

// Docoders & Encoders
func decodeGetUsersRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	var req getUsersRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func decodeAddUsersRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	var req addUserRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}
