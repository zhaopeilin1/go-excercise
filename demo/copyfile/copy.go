package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	src := "赵培淋-5年后端开发.pdf"
	dest := "赵培淋-5年后端开发22.pdf"
	bytes, err := copyFile(src, dest)
	fmt.Println(bytes, err)
}

func copyFile(src, dest string) (int, error) {
	file1, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(dest, os.O_CREATE|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()

	bs := make([]byte, 1024, 1024)
	n := -1
	total := 0
	for {
		n, err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("copy completed")
			break
		} else {
			total += n
			file2.Write(bs[:n])
		}
	}
	return total, nil
}
