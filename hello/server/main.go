package main

import (
	"Grpc-Demo/proto"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedUserServer
}

func (s *server) UserReg(ctx context.Context, req *proto.UserRegRequest) (*proto.UserRegResponse, error) {
	fmt.Println("request:", req.Nickname)
	fmt.Println("request:", req)
	//todo 业务实现---begin

	//todo 业务实现---end

	return &proto.UserRegResponse{Id: 1}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8001")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	proto.RegisterUserServer(s, &server{})
	//reflection.Register(s)

	defer func() {
		s.Stop()
		listen.Close()
	}()

	fmt.Println("Serving 8001...")
	err = s.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
