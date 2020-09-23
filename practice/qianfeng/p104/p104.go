package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("aa.txt")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("create file:", file)

	file1, err1 := os.Open("bb.txt")
	if err1 != nil {
		fmt.Println("err:", err1)
	}
	fmt.Println("create file:", file1)

}
