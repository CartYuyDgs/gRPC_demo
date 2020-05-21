package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	protofile_hello "gRPC_demo/examples/protofiles"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	var serverHost = "127.0.0.1:50001"
	conn, err := grpc.Dial(serverHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connect err %v", err)
	}
	defer conn.Close()

	log.Println("connect success!")

	client := protofile_hello.NewHelloServiceClient(conn)
	for {
		fmt.Println("请输入名字：")
		name, err := GetInput()
		if err != nil {
			log.Fatalf("input error %v!", err)
			continue
		}
		fmt.Println("请输入消息:")
		mes, err := GetInput()
		if err != nil {
			log.Fatalf("mes error %v!", err)
			continue
		}
		resp, err := client.SayHello(context.Background(), &protofile_hello.HelloRequest{Name: name, Message: mes})
		if err != nil {
			log.Fatalf("Error sayhello err %v", err)
		}

		fmt.Println("Server Response: name:", resp.Req.Name, " massage:", resp.Req.Message)
	}

}

func GetInput() (string, error) {
	var input string

	in := bufio.NewScanner(os.Stdin)
	if in.Scan() {
		input = in.Text()
		return input, nil
	}
	return "", errors.New("err input ")
}
