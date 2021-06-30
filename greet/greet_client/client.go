package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"

	"../greetpb"
)

func main() {
	fmt.Println("Greet Client")
	
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	// fmt.Printf("Created client: %f", c)

	doUnary(c)

	doServerStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Greet Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting {
			FirstName: "Calvin",
			Lastname: "Colton",
		}
	}
	res, err := c.Greet(context.Background(), in req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting {
			FirstName: "Calvin",
			Lastname: "Colton",
		}
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error whil ereading stream %v:", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}