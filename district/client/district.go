package client

import (
	"context"

	grpcClient "github.com/region-apps/grpc"
	pb "github.com/region-apps/grpc/district/proto"
)

func (c *Client) GetAllDistricts(ctx context.Context) (*pb.GetAllDistrictsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.DistrictClient.GetAllDistricts(ctx, nil)
}

func (c *Client) GetDistrictByID(ctx context.Context, req *pb.GetDistrictByIDRequest) (*pb.GetDistrictByIDResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.DistrictClient.GetDistrictByID(ctx, req)
}

func (c *Client) GetDistrictsByIDs(ctx context.Context, req *pb.GetDistrictsByIDsRequest) (*pb.GetDistrictsByIDsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.DistrictClient.GetDistrictsByIDs(ctx, req)
}
