package main

import (
	"Grpc-Demo/proto"
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func main() {

	var serviceHost = "127.0.0.1:8001"

	conn, err := grpc.Dial(serviceHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	client := proto.NewUserClient(conn)
	rsp, err := client.UserReg(context.TODO(), &proto.UserRegRequest{
		Nickname: "parker",
		Age:      1,
		Sex:      2,
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.Age)

	fmt.Println("按回车键退出程序...")
	in := bufio.NewReader(os.Stdin)
	_, _, _ = in.ReadLine()
}
