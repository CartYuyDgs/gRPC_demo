package main

import (
	"context"
	"fmt"
	protofile_hello "gRPC_demo/examples/protofiles"
	"google.golang.org/grpc"
	"log"
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
	resp, err := client.SayHello(context.Background(), &protofile_hello.HelloRequest{Name: "xiaoming"})
	if err != nil {
		log.Fatalf("Error sayhello err %v", err)
	}

	fmt.Println("name:", resp.Req.Name, " massage:", resp.Req.Message)
}
