package service

import (
	carrpc "carRider-srv/proto/rpc"
	"context"
)

type CartServer struct {
	carrpc.UnimplementedCartServer
}

func (s *CartServer) Demo(_ context.Context, in *carrpc.DemoReq) (*carrpc.DemoResp, error) {
	return &carrpc.DemoResp{
		Demo: in.Name,
	}, nil
}
