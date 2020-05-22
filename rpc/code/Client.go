package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {

	var (
		num  = 5.0
		resp float32
	)

	client, err := rpc.DialHTTP("tcp", "localhost:10000")
	if err != nil {
		log.Fatalf("Error %v", err)
		return
	}

	err = client.Call("MathUtil.Calculate", num, &resp)
	if err != nil {
		log.Fatalf("Error %v", err)
		return
	}

	fmt.Println(resp)
}
