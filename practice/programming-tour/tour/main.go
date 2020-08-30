package main

import (
	"log"
	"programming-tour/tour/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal("cmd.Execute err:")
	}
}
