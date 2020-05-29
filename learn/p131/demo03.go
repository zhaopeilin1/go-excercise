// demo03
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name   string
	Age    int
	School string
}

func main() {
	s1 := Student{"王二狗", 10, "千峰教育"}
	//通过反射修改对象的值 ，前提是可建起修改
	fmt.Printf("%T\n", s1)
	p1 := &s1
	fmt.Printf("%T\n", p1)

	//改变数值
	value := reflect.ValueOf(p1)
	newValue := value.Elem()
	fmt.Println(newValue.CanSet())
	//根据名称获取字段，反射。然后修改字段的值。
	fieldValue := newValue.FieldByName("Name")
	fieldValue.SetString("小明")
	fmt.Println(s1)
}
