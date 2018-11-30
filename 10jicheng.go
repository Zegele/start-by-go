package main

import (
	"fmt"
)

type human struct {
	Age int
}

type teacher struct {
	human
	Name string
	Age  int
}

type student struct {
	human
	Name string
	Age  int
}

func main() {
	a := teacher{
		Name:  "Joson",
		Age:   19,
		human: human{Age: 1},
		//在teacher和student中使用了另一种匿名字段，
		//系统默认了human就是该字段的名称,虽然human也是其中的嵌套结构
	}
	b := student{Name: "Holo", Age: 18, human: human{Age: 0}} //这样写好点
	fmt.Print(a, b)

	a.Name = "Liu"
	a.Age = 29
	a.human.Age = 100

	fmt.Print(a, b)

	a.Age = 1000 //如果同名，则优先本结构，不考虑嵌套结构。
	fmt.Print(a, b)

}
