package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"grpc-demo/gen/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.Add(ctx, &proto.AddRequest{A: 3, B: 5})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}

	fmt.Printf("Result: %d\n", resp.Result)
}
