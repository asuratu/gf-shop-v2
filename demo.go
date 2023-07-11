package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func parse() {
	m := make(map[string]*student)
	g := make(map[string]int)
	students := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for i, stu := range students {
		// 这里是关键，stu是students的一个副本, 在内存中, stu的地址是不变的
		// m[stu.Name] = &stu
		m[stu.Name] = &students[i]
		g[stu.Name] = stu.Age
		stu.Age = 0
	}

	fmt.Println(m)
	fmt.Println(students)
}

func main() {
	parse()
}
