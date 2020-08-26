package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func main() {
	pase_student()
}

func pase_student() {
	m := make(map[string]*student)

	stus := []student{
		student{Name: "a", Age: 1},
		student{Name: "b", Age: 2},
		student{Name: "c", Age: 3},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	//fmt.Println(m)
}
