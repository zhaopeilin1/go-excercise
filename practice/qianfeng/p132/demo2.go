// demo2
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//反射调用函数
	f1 := fun1
	value := reflect.ValueOf(f1)
	fmt.Printf("kind,%s,type:%s\n", value.Kind(), value.Type())
	value2 := reflect.ValueOf(fun2)
	fmt.Printf("2kind,%s,type:%s\n", value2.Kind(), value2.Type())
	value3 := reflect.ValueOf(fun3)
	fmt.Printf("3kind,%s,type:%s\n", value3.Kind(), value3.Type())
	value4 := value3.Call([]reflect.Value{reflect.ValueOf(1), reflect.ValueOf("aa")})
	fmt.Printf("%T\n", value4)
	fmt.Println(value4)
	fmt.Printf("kind:%s,type:%s\n", value4[0].Kind(), value4[0].Type())
	//结果变为实际类型
	s := value4[0].Interface().(string)
	fmt.Println(s)
}

func fun1() {
	fmt.Println("无参函数fun1")
}
func fun2(i int, s string) {
	fmt.Println("有参函数fun2")
}

func fun3(i int, s string) string {
	fmt.Println("有参有返回函数fun3")
	return strconv.Itoa(i) + s
}
