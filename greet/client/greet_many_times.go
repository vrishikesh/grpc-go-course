package main

import (
	"context"
	pb "grpc-go-course/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Printf("doGreetManyTimes function was invoked")

	req := &pb.GreetRequest{
		FirstName: "Rishi",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
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

		log.Printf("GreetManyTimes: %s", msg.Result)
	}
}
