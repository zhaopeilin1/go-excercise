package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	readFile2("read.txt")
	writeFile("write.txt", []byte("hello\n world!"))

}

func readFile(filename string) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(b)
	str := string(b)
	fmt.Println(str)
}
func readFile2(filename string) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(bs)
		str := string(bs)
		fmt.Println(str)
	}
}

func writeFile(fileName string, bs []byte) {
	err := ioutil.WriteFile(fileName, bs, 775)
	if err != nil {
		fmt.Println(err)
	}

}
