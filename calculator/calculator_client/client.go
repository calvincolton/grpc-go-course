package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"

	"../calculatorpb"
)

func main() {
	fmt.Println("Calculator Client")
	
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	// fmt.Printf("Created client: %f", c)

	doUnary(c)

	doServerStreaming(c)

	doClientStreaming(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Sum Unary RPC...")
	req := &calculatorpb.SumRequest{
		FirstNumber: 5,
		SecondNumber: 40,
	}

	res, err := c.Sum(context.Background(), in req)
	if err != nil {
		log.FatalF("error while calling sum RPC: %v", err)
	}

	log.Printf("Response from Sum: %v", res.SumResult)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a PrimeNumberDecomposition Server Streaming RPC...")
	req := &calculatorpb.PrimeNumberDecomposition{
		Number: 12,
	}

	stream, err := c.PrimeNumberDecomposition(context.Background(), in req)
	if err != nil {
		log.FatalF("error while calling PrimeNumberDecomposition RPC: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened %v", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}
	log.Printf("Response from Sum: %v", res.SumResult)
}

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a ComputeAverage Client Streaming RPC...")

	stream, err := c.ComputeAverage(context.Bacckground())
	if err != nil {
		log.Fatalf("Error while opening stream: %v", err)
	}

	numbers := []int32{3, 5, 8, 9, 18, 20, 24, 57, 61, 77}

	for _, number := range numbers {
		fmt.Printf("Sending number: %v\n", number)
		stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: number,
		}),
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}

	fmt.Printf("The average is: %v", res.GetAverage())
}