package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	args := os.Args[1:]
	if len(args) == 0 {
		countsFile(os.Stdin, counts)
	} else {
		for _, arg := range args {
			file, err := os.Open(arg)
			//要关闭文件
			defer file.Close()
			if err != nil {
				fmt.Println(" open file error", arg, err)
				continue
			}
			countsFile(file, counts)
			file.Close()

		}

	}
	outputMap(counts)

}
func outputMap(counts map[string]int) {
	for content, counts := range counts {
		if counts > 1 {
			fmt.Println(counts, content)
		}
	}

}

func countsFile(file *os.File, counts map[string]int) {
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		counts[scan.Text()]++
	}

}
