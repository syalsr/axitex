package main

import (
	"axitex/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	addr string
	num1 float32
	num2 float32
)

func init() {
	addr = os.Args[1]
	
	tmp1, _ := strconv.ParseFloat(os.Args[2], 32)
	num1 = float32(tmp1)

	tmp2, _ := strconv.ParseFloat(os.Args[3], 32)
	num2 = float32(tmp2)
}

func main() {
	conn, err := grpc.Dial(
		addr, grpc.WithTransportCredentials(
			insecure.
				NewCredentials(),
		),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewDatabusServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	r, err := c.Send(ctx, &proto.SendRequest{Prm1: num1, Prm2: num2})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("Result is %f\n", r.Result)
}
