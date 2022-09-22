package main

import (
	"axitex/proto"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type DatabusService struct {
	proto.UnimplementedDatabusServiceServer
}

const (
	add string = "add"
	sub        = "sub"
	mul        = "mul"
	div        = "div"
)

func calculate(num1 float32, num2 float32, op string) (float32, error) {
	switch op {
	case add:
		res := num1 + num2
		fmt.Printf("%f + %f=%f\n", num1, num2, res)
		return res, nil
	case sub:
		res := num1 - num2
		fmt.Printf("%f - %f=%f\n", num1, num2, res)
		return res, nil
	case mul:
		res := num1 * num2
		fmt.Printf("%f * %f=%f\n", num1, num2, res)
		return res, nil
	case div:
		res := num1 / num2
		fmt.Printf("%f \\ %f=%f\n", num1, num2, res)
		return res, nil
	}
	return 0, errors.New("unknown type of operation")
}

func (d *DatabusService) Send(ctx context.Context, in *proto.SendRequest) (
	*proto.SendResponse,
	error,
) {
	res, err := calculate(in.Prm1, in.Prm2, operation)
	if err != nil {
		log.Fatal("error")
		return &proto.SendResponse{}, nil
	}
	return &proto.SendResponse{Result: res}, nil
}

var (
	port      string
	operation string
)

func init() {
	port = os.Args[1]
	operation = os.Args[2]
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterDatabusServiceServer(s, &DatabusService{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
