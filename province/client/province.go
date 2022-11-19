package client

import (
	"context"

	grpcClient "github.com/region-apps/grpc"
	pb "github.com/region-apps/grpc/province/proto"
)

func (c *Client) GetAllProvinces(ctx context.Context) (*pb.GetAllProvincesResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.ProvinceClient.GetAllProvinces(ctx, nil)
}

func (c *Client) GetProvinceByID(ctx context.Context, req *pb.GetProvinceByIDRequest) (*pb.GetProvinceByIDResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.ProvinceClient.GetProvinceByID(ctx, req)
}

func (c *Client) GetProvincesByIDs(ctx context.Context, req *pb.GetProvincesByIDsRequest) (*pb.GetProvincesByIDsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.ProvinceClient.GetProvincesByIDs(ctx, req)
}
