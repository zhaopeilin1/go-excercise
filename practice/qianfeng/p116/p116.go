package main

import (
	"fmt"
	"runtime"
)

func init() {

}

func main() {
	fmt.Println(runtime.GOROOT(), runtime.GOOS, runtime.NumCPU())
}
