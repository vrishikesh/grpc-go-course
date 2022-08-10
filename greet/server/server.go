package main

import pb "grpc-go-course/greet/proto"

type Server struct {
	pb.GreetServiceServer
}
