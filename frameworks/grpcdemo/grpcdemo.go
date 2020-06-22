package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type MyType struct {
}

func (t MyType) Add(a int64, b *int64) error {
	fmt.Println("rpc call")
	*b = a / 2
	return nil
}

func main() {
	t := new(MyType)
	rpc.Register(t)
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	http.Serve(listen, nil)

}
