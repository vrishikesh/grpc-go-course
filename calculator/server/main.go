package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpc-go-course/calculator/proto"
)

var addr string = "0.0.0.0:50052"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("could listen to %s: %v", addr, err)
	}

	log.Printf("listening on %s", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("could not serve: %v", err)
	}
}
