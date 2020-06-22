package main

import "fmt"

func main() {
	slice := make([]int, 0)
	slice = append(slice, 22)
	for index, value := range slice {
		fmt.Println(index, value)
	}
}
