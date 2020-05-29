package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.23
	fmt.Println(num)

	//需要操作指针
	pointValue := reflect.ValueOf(&num)
	//需要是指针，否则 报错 获取 指针所对应的对象的的值。
	v2 := pointValue.Elem()
	fmt.Println("类型：", v2.Type()) //float64
	fmt.Println(v2.CanSet())      //true
	//指针拿到的值，可以修改
	v2.SetFloat(3.14)
	fmt.Println(num)

	value := reflect.ValueOf(num)
	// value.SetFloat(1.11)//panic
	value.CanSet() //false
	value.Elem()   //panic

}
