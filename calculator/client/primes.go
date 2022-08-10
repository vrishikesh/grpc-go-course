package main

import (
	"context"
	pb "grpc-go-course/calculator/proto"
	"io"
	"log"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Printf("doPrimes function was invoked")

	stream, err := c.Primes(context.Background(), &pb.PrimesRequest{
		N: 120,
	})
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while reading the stream: %v", err)
		}

		log.Printf("Primes: %d", msg.Result)
	}
}
