package main

import pb "grpc-go-course/calculator/proto"

type Server struct {
	pb.CalculatorServiceServer
}
