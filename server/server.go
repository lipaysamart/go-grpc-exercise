package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/lipaysamart/go-grpc-exercise/server/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Reply: "Hello " + in.Name + " Age " + fmt.Sprintf("%d", in.Age)}, nil
}

func main() {
	conn, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gsrv := grpc.NewServer()
	pb.RegisterGreeterServer(gsrv, &server{})

	err = gsrv.Serve(conn)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
