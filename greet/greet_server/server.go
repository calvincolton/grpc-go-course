package main

import (
	"fmt"
	"log"
	"net"

	"../../greet/greetpb"

	"google.golang.org/grpc"
	// "github.com/simplestpath/grpc-go-course/greet/greetpb"
)

type server struct {}

// type server struct {
// 	greetpb.UnimplementedGreetServiceServer
}

type GreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	// greetpb.RegisterGreetSwerviceServer(s *grpc.Server, srv greetpb.GreetServiceServer)
	// greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Server(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}