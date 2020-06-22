package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err)
	}
	a := 32
	var b *int64
	client.Call("MyType.Add", a, b)
	fmt.Println(b)
}
