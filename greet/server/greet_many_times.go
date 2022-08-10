package main

import (
	"fmt"
	pb "grpc-go-course/greet/proto"
	"log"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked: %v", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.GetFirstName(), i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
