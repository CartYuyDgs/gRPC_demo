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
	//同步
	//err = client.Call("MathUtil.Calculate", num, &resp)
	//if err != nil {
	//	log.Fatalf("Error %v", err)
	//	return
	//}

	//异步
	asyncCall := client.Go("MathUtil.Calculate", num, &resp, nil)
	replayDone := <-asyncCall.Done
	fmt.Println(replayDone)

	fmt.Println(resp)
}
