package main

import (
	"log"
	"math"
	"net"
	"net/http"
	"net/rpc"
)

type MathUtil struct {
}

func (ma *MathUtil) Calculate(req float32, resp *float32) error {
	*resp = math.Pi * req * req
	return nil
}

func main() {
	math := new(MathUtil)

	err := rpc.Register(math)
	if err != nil {
		log.Fatalf("Error %v", err)
		return
	}

	rpc.HandleHTTP()

	listen, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatalf("Error %v", err)
		return
	}

	http.Serve(listen, nil)
}
