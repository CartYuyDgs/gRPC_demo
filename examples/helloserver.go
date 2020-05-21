package main

import (
	"fmt"
	"gRPC_demo/examples/impl"
	protofile_hello "gRPC_demo/examples/protofiles"
	"google.golang.org/grpc"
	"log"
	"net"
)

var port = 50001

//启动grpc服务
func main() {
	lisn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v ", err)
	}
	log.Printf("listen: %d\n", port)
	grpcServer := grpc.NewServer()
	protofile_hello.RegisterHelloServiceServer(grpcServer, &impl.HelloService{})
	grpcServer.Serve(lisn)
}
