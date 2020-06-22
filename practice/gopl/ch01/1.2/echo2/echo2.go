package main

import (
	"fmt"
	"os"
)

func main() {
	s, seq := "", ""
	for _, arg := range os.Args[0:] {
		s += seq + arg
		seq = ":"
	}
	fmt.Println(s)
}
