package main

import (
	"flag"
	"fmt"
	pb "github.com/potix2/grpc-with-bazel/hello/hello"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

var (
	addr = flag.String("addr", "localhost:12345", "host:port")
)

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*addr, opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)
	response, err := client.SayHello(context.Background(), &pb.HelloRequest{
		Greeting: "world",
	})
	if err != nil {
		log.Fatalf("failed to call SayHello() = %v, %v", response, err)
	}
	fmt.Println(response.Reply)
}
