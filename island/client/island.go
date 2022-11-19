package client

import (
	"context"

	grpcClient "github.com/region-apps/grpc"
	pb "github.com/region-apps/grpc/island/proto"
)

func (c *Client) GetAllIslands(ctx context.Context) (*pb.GetAllIslandsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.IslandClient.GetAllIslands(ctx, nil)
}

func (c *Client) GetIslandByID(ctx context.Context, req *pb.GetIslandByIDRequest) (*pb.GetIslandByIDResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.IslandClient.GetIslandByID(ctx, req)
}

func (c *Client) GetIslandsByIDs(ctx context.Context, req *pb.GetIslandsByIDsRequest) (*pb.GetIslandsByIDsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.IslandClient.GetIslandsByIDs(ctx, req)
}
