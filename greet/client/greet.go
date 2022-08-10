package main

import (
	"context"
	pb "grpc-go-course/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Printf("doGreet was invoked")
	r, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Rishi",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.Result)
}
