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