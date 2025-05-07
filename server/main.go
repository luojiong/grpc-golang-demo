package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpc-demo/gen/proto"

	"google.golang.org/grpc"
)

type calculatorServer struct {
	proto.UnimplementedCalculatorServer
}

func (s *calculatorServer) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	result := req.A + req.B
	return &proto.AddResponse{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterCalculatorServer(grpcServer, &calculatorServer{})
	fmt.Println("Server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
