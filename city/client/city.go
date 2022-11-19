package client

import (
	"context"

	grpcClient "github.com/region-apps/grpc"
	pb "github.com/region-apps/grpc/city/proto"
)

func (c *Client) GetAllCities(ctx context.Context) (*pb.GetAllCitiesResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.CityClient.GetAllCities(ctx, nil)
}

func (c *Client) GetCityByID(ctx context.Context, req *pb.GetCityByIDRequest) (*pb.GetCityByIDResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.CityClient.GetCityByID(ctx, req)
}

func (c *Client) GetCitiesByIDs(ctx context.Context, req *pb.GetCitiesByIDsRequest) (*pb.GetCitiesByIDsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.CityClient.GetCitiesByIDs(ctx, req)
}
