package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-go-course/calculator/proto"
)

var addr string = "0.0.0.0:50052"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could could connect to calculator service: %v", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)
	// doSum(c)
	doPrimes(c)
}
