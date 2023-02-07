package main

import (
	"context"
	"log"
	"net"

	pb "github.com/lzzzzl/gogo-learn/grpc/grpc-helloworld/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{})
	log.Println("Server is running on port 0.0.0.0:50051...")
	log.Fatal(s.Serve(lis))
}
