package main

//client.go

import (
	"log"
	"os"

	pb "gotest/hello/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "139.199.200.93:50051"
	defaultName = "world"
)



func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
