package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "192.168.43.175:8081")
	if err != nil {
		log.Fatal(err)
		return
	}
	var aa *int64
	t1 := time.Now()
	client.Call("MyType.Add", 77, &aa)
	//t2 := time.Now()
	fmt.Println(*aa)
	fmt.Println(time.Since(t1))

	asnyCall := client.Go("MyType.Add", 22, &aa, nil)
	replayDone := <-asnyCall.Done
	fmt.Println(replayDone)
	fmt.Println(*aa)
}
