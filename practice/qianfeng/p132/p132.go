package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) Say(msg string) {
	fmt.Println("hello", msg)
}
func (p Person) PrintInfo() {
	fmt.Printf("姓名%s,年龄%d,性别 %s\n", p.Name, p.Age, p.Sex)
}

func main() {
	//reflect来进行方法的调用。
	p1 := Person{"王二狗", 10, "男"}
	value := reflect.ValueOf(p1)
	fmt.Printf("kind:%s,type:%s\n", value.Kind(), value.Type())
	method1 := value.MethodByName("PrintInfo")
	fmt.Printf("kind:%s,type:%s\n", method1.Kind(), method1.Type())
	method1.Call(nil) //如果没有参数直接传nil或空切片
	method1.Call(make([]reflect.Value, 0))

	method2 := value.MethodByName("Say")
	fmt.Printf("kind:%s,type:%s\n", method2.Kind(), method2.Type())
	args2 := []reflect.Value{reflect.ValueOf("反射的值")}
	method2.Call(args2)

}
