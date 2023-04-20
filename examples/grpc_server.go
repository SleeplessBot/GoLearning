package examples

import (
	"context"
	"fmt"
	"log"
	"net"

	"golearning/examples/pb"

	"google.golang.org/grpc"
)

type GreetServer struct {
	pb.UnimplementedGreetServer
}

func RunGrpcServer() {
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGreetServer(grpcServer, &GreetServer{})

	log.Printf("GRPC server listening on %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *GreetServer) SayHi(ctx context.Context, req *pb.SayHiRequest) (*pb.SayHiReply, error) {
	fmt.Printf("Get greet from %s\n", req.Name)
	res := &pb.SayHiReply{
		Message: "Hi, " + req.Name + "!",
	}
	return res, nil
}
