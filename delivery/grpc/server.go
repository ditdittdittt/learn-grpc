package deliverygrpc

import (
	"context"

	pb "github.com/ditdittdittt/learn-grpc/productproto"
)

type productServer struct {
	pb.UnimplementedProductServer
}

func (s *productServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{Result: req.GetName()}, nil
}

func NewProductServer() *productServer {
	return &productServer{pb.UnimplementedProductServer{}}
}
