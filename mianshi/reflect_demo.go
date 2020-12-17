package main

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

func main() {
	//fmt.Println("ok")

	person := Person{
		Name: "aa",
	}
	//fmt.Println(person)
	s := Student{Age: 11, Person: person}
	//s := &Student{Age: 11,}
	m := structToMap(s)
	fmt.Println(m)

	te:= Teacher{Name: "tt",Age: 11}

	m2:=structToMap2(&te)

	fmt.Println(m2)
}

type Teacher struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type Person struct {
	Name string `json:"name"`
}

type Student struct {
	Age int `json:"age"`
	Person
}


func structToMap2(i interface{}) url.Values {
	values := url.Values{}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		// You ca use tags here...
		tag := typ.Field(i).Tag.Get("json")
		fmt.Println("tag"+fmt.Sprint(i)+":" +tag)
		// Convert each type into a string for the url.Values string map
		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
			//case interface{}:
			//	subValues:=structToMap(f)
			//	for k,v1 := range subValues{
			//		values[k] = v1
			//	}
		}
		if tag!="" {
			values.Set(tag, v)
		}

	}
	return values
}

func structToMap(i interface{}) url.Values {
	values := url.Values{}
	getValue := reflect.ValueOf(i)
	getType := reflect.TypeOf(i)
	for i := 0; i < getValue.NumField(); i++ {
		f := getValue.Field(i)
		// You ca use tags here...
		tag := getType.Field(i).Tag.Get("json")
		fmt.Println("tag" + fmt.Sprint(i) + ":" + tag)
		// Convert each type into a string for the url.Values string map
		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		case interface{}:
			//嵌套结构体处理
			subValues := structToMap(f.Interface())
			for k, v1 := range subValues {
				values[k] = v1
			}
		}
		if tag!="" {
			//嵌套字段，tag为空的，不加入map中
			values.Set(tag, v)
		}
	}
	return values
}
