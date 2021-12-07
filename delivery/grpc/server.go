package deliverygrpc

import (
	"context"

	pb "github.com/ditdittdittt/learn-grpc/productproto"
	"github.com/ditdittdittt/learn-grpc/usecase"
)

type server struct {
	productUsecase usecase.ProductUsecase
	pb.UnimplementedProductServiceServer
}

func (s *server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	reqName := req.GetName()

	respName, err := s.productUsecase.Create(ctx, reqName)
	if err != nil {
		return nil, err
	}

	resp := &pb.CreateResponse{
		ID:           respName,
		ResponseCode: 200,
	}

	return resp, nil
}

func (s *server) Read(ctx context.Context, req *pb.ReadRequest) (*pb.ReadResponse, error) {
	productMap, err := s.productUsecase.Read(ctx)
	if err != nil {
		return nil, err
	}

	resp := &pb.ReadResponse{}
	for id, name := range productMap {
		resp.Products = append(resp.Products, &pb.Product{
			ID:   id,
			Name: name,
		})
	}
	resp.ResponseCode = 200

	return resp, nil
}

func (s *server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	reqID := req.GetID()
	reqName := req.GetName()

	respID, respName, err := s.productUsecase.Update(ctx, reqID, reqName)
	if err != nil {
		return nil, err
	}

	resp := &pb.UpdateResponse{
		ID:           respID,
		Name:         respName,
		ResponseCode: 200,
	}

	return resp, nil
}

func (s *server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	reqID := req.GetID()

	err := s.productUsecase.Delete(ctx, reqID)
	if err != nil {
		return nil, err
	}

	resp := &pb.DeleteResponse{
		ResponseCode: 200,
	}

	return resp, nil
}

func NewProductServiceServer(productUsecase usecase.ProductUsecase) *server {
	return &server{
		productUsecase:                    productUsecase,
		UnimplementedProductServiceServer: pb.UnimplementedProductServiceServer{},
	}
}
