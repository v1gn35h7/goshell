package http

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/v1gn35h7/goshell/pkg/goshell"
	"github.com/v1gn35h7/goshell/server/service"
)

// Endpoints creators
func MakeGetAssetsEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		//req := request.(getAssetsdRequest)
		res, err := srvc.GetAssets()

		if err != nil {
			return getAssetsResponse{List: make([]goshell.Asset, 0), Count: 0}, err
		}

		data := make([]goshell.Asset, 0)
		for _, v := range res {
			data = append(data, *v)
		}

		return getAssetsResponse{List: data, Count: int64(len(res))}, nil
	}
}
