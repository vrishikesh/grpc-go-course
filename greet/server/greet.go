package main

import (
	"context"
	pb "grpc-go-course/greet/proto"
	"log"
)

func (*Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("greet function invoked with %v", in)

	return &pb.GreetResponse{
		Result: "Hello " + in.GetFirstName(),
	}, nil
}
