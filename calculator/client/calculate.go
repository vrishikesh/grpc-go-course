package main

import (
	"context"
	pb "grpc-go-course/calculator/proto"
	"log"
)

func doSum(c pb.CalculatorServiceClient) {
	nums := []int32{1, 2, 3}
	r, err := c.Sum(context.Background(), &pb.SumRequest{
		Numbers: nums,
	})
	if err != nil {
		log.Fatalf("could not sum: %v", err)
	}

	log.Printf("Sum of %v is %d", nums, r.Result)
}
