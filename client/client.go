package main

import (
	"context"
	"flag"
	"log"

	"github.com/lipaysamart/go-grpc-exercise/client/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", "world", "Name to greet")
)

func main() {
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name, Age: 20})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetReply())

}
