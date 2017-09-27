package main

import (
	"flag"
	pb "github.com/potix2/grpc-with-bazel/hello/hello"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	addr = flag.String("addr", "localhost:12345", "host:port")
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Reply: req.Greeting,
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listend: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}
