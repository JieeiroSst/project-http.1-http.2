package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/JieeiroSst/project-grpc/proto"
	"google.golang.org/grpc"
)

type server struct {}

func (s *server) Tall(ctx context.Context, in *proto.Message) (*proto.Message, error) {
	return in, nil
}

func (s *server) Sum(ctx context.Context, in *proto.SumRequest) (*proto.SumResponse, error) {
	return &proto.SumResponse{
		Result: in.Num1 + in.Num2,
	}, nil
}

func (s *server) Subtraction(ctx context.Context, in *proto.SumRequest) (*proto.SumResponse, error) {
	return &proto.SumResponse{
		Result: in.Num1 - in.Num2,
	}, nil
}

func (s *server) Multiplication(ctx context.Context, in *proto.SumRequest) (*proto.SumResponse, error) {
	return &proto.SumResponse{
		Result: in.Num1 * in.Num2,
	}, nil
}

func (s *server) Division(ctx context.Context, in *proto.SumRequest) (*proto.SumResponse, error) {
	if in.Num2 < 0 {
		return &proto.SumResponse{}, errors.New("num2 == 0")
	}
	return &proto.SumResponse{
		Result: in.Num1 + in.Num2,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatalf("err while create listen  %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterServiceServer(s, &server{})

	fmt.Printf("server is running...")

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}

