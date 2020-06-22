package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		反射
	*/
	var x Person
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))
	fmt.Println("------")
	v := reflect.ValueOf(x)
	fmt.Println("kind", v.Kind(), "type", v.Type(), "value", v)

	var num float64 = 1.23
	value := reflect.ValueOf(num)
	convertValue := value.Interface().(float64)
	fmt.Println(convertValue)
	pointer := reflect.ValueOf(&num)
	convertPointer := pointer.Interface().(*float64)
	fmt.Println(convertPointer)
}

type Person struct {
	name string
	age  int
}
