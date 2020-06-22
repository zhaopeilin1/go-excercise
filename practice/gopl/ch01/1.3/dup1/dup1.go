package main

import (
        "fmt"
        "os"
        "bufio"
)

func main(){
        countsMap := make(map[string]int)
        input := bufio.NewScanner(os.Stdin)
        for input.Scan(){
                countsMap[input.Text()]++
        }
        for line,n:=range countsMap {
                if n>1 {
                        fmt.Printf("%d \t %s \n",n ,line)
                }
        }
}

