package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	deliverygrpc "github.com/ditdittdittt/learn-grpc/delivery/grpc"
	pb "github.com/ditdittdittt/learn-grpc/productproto"
	"github.com/ditdittdittt/learn-grpc/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Setup Usecase
	productMap := map[string]string{}
	productUsecase := usecase.NewProductUsecase(productMap)
	productGRPC := deliverygrpc.NewProductServiceServer(productUsecase)

	// Setup Server
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterProductServiceServer(s, productGRPC)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
