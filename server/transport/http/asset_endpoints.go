package http

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/v1gn35h7/goshell/server/service"
)

// Endpoints creators
func MakeGetAssetsEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		//req := request.(getAssetsdRequest)
		res, err := srvc.GetAssets()

		if err != nil {
			return getAssetsResponse{List: make([]string, 0), Count: 0}, err
		}

		return getAssetsResponse{List: res, Count: int64(len(res))}, nil
	}
}
