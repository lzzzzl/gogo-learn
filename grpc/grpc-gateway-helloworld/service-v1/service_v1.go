/**
简单版服务
**/

package main

import (
	"context"
	"log"
	"net"

	pb "github.com/lzzzzl/gogo-learn/grpc/grpc-gateway-helloworld/proto/helloworld"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln("Failed to listen: ", err)
	}

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &Server{})
	log.Println("Server is running on port 0.0.0.0:8081...")
	log.Fatal(s.Serve(lis))
}
