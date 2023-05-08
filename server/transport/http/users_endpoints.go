package http

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/goShell/server/service"
)

// Endpoints creators

func MakeGetUsersEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		//req := request.(getAssetsdRequest)
		res, err := srvc.GetUsers()

		if err != nil {
			return getAssetsResponse{List: make([]string, 0), Count: 0}, err
		}

		return getAssetsResponse{List: res, Count: int64(len(res))}, nil
	}
}

func MakeAddUsersEndpoint(srvc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(addUserRequest)
		res, err := srvc.AddUser(req.User)

		if err != nil {
			return addUserResponse{Status: "Failed to add"}, err
		}

		return addUserResponse{Status: res}, nil
	}
}
