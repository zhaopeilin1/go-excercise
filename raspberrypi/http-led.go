cat httpled.go 
package main

import (
	"fmt"
	"github.com/nathan-osman/go-rpigpio"
	"log"
	"net/http"
	"time"
)

var n int = 14
var status int

func main() {
	p, err := rpi.OpenPin(n, rpi.OUT)
	if err != nil {
		panic(err)
	}
	defer p.Close()

	//set high
	fmt.Println("start gpio test,pin out:", n)

	go ledAction(p)
	http.HandleFunc("/up", up)
	http.HandleFunc("/down", down)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func up(w http.ResponseWriter, r *http.Request) {
	//使led灯闪烁
	n = 1
	fmt.Fprintf(w,"开启闪光灯，气氛搞起来")
}
func down(w http.ResponseWriter, r *http.Request) {
	//使led灯熄灭
	n = 0
	fmt.Fprintf(w,"关闭闪光灯")
}
func ledAction(p *rpi.Pin) {
	for {
		time.Sleep(time.Millisecond * 100)
		if n == 1 {
			//set hight led up
			p.Write(rpi.HIGH)
			p.Write(rpi.LOW)
		}
	}
}
