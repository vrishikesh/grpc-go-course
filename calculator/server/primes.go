package main

import (
	"log"

	pb "grpc-go-course/calculator/proto"
)

func (*Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function invoked with %v", in)

	var N int32 = in.N
	var k int32 = 2
	for N > 1 {
		if N%k == 0 {
			log.Printf("factor found %d", k)
			stream.Send(&pb.PrimesResponse{
				Result: k,
			})
			N = N / k
		} else {
			k = k + 1
		}
	}

	return nil
}
