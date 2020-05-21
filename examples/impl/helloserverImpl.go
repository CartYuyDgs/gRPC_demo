package impl

import (
	"context"
	protofile_hello "gRPC_demo/examples/protofiles"
)

type HelloService struct {
}

//实现Rpc接口
func (h *HelloService) SayHello(ctx context.Context, request *protofile_hello.HelloRequest) (*protofile_hello.HelloResponse, error) {
	return &protofile_hello.HelloResponse{
		Req: &protofile_hello.Request{
			Name:    request.Name,
			Message: request.Message,
			IsTrue:  true,
		},
	}, nil
}
