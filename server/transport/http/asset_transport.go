package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/v1gn35h7/goshell/pkg/goshell"
)

// Asset service contracts
type getAssetsRequest struct {
	Query string `json:"query"`
}

type getAssetsResponse struct {
	List  []goshell.Asset `json:"list"`
	Count int64           `json:"count"`
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
	req := getAssetsRequest{
		Query: request.URL.Query().Get("hostId"),
	}

	return req, nil
}

func encodeGetAssetsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
