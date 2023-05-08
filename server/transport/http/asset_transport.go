package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// Asset service contracts
type getAssetsRequest struct {
	Query string `json:"query"`
}

type getAssetsResponse struct {
	List  []string `json:"list"`
	Count int64    `json:"count"`
}

// Scaffloding endpoints to transport
func getAssetsEnpointTransport(endpoint endpoint.Endpoint) http.Handler {
	return httptransport.NewServer(
		endpoint,
		decodeGetAssetsRequest,
		encodeGetAssetsResponse,
	)
}

// Request utilities
func decodeGetAssetsRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	var req getAssetsRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func encodeGetAssetsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
