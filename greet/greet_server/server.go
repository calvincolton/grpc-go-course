package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"../greetpb"
)

type server struct {}

// type server struct {
// 	greetpb.UnimplementedGreetServiceServer
// }

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstNam()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func(*server) GreetManyTimes(req *greetpb.GreetManytimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i< 10; i++; {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res.&greetpb.GreetManyTimesReponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

type GreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func main() {
	fmt.Println("Greet Server")

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