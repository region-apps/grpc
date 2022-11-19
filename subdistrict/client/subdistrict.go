package client

import (
	"context"

	grpcClient "github.com/region-apps/grpc"
	pb "github.com/region-apps/grpc/subdistrict/proto"
)

func (c *Client) GetAllSubdistricts(ctx context.Context) (*pb.GetAllSubdistrictsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.SubdistrictClient.GetAllSubdistricts(ctx, nil)
}

func (c *Client) GetSubdistrictByID(ctx context.Context, req *pb.GetSubdistrictByIDRequest) (*pb.GetSubdistrictByIDResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.SubdistrictClient.GetSubdistrictByID(ctx, req)
}

func (c *Client) GetSubdistrictsByIDs(ctx context.Context, req *pb.GetSubdistrictsByIDsRequest) (*pb.GetSubdistrictsByIDsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.SubdistrictClient.GetSubdistrictsByIDs(ctx, req)
}
