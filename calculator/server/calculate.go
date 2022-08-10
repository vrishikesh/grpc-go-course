package main

import (
	"context"
	"log"

	pb "grpc-go-course/calculator/proto"
)

func (*Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("sum function invoked with %v", in.Numbers)

	var sum int32
	for _, i := range in.Numbers {
		sum += i
	}

	return &pb.SumResponse{
		Result: sum,
	}, nil
}
