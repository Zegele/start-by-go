package main

import (
	"fmt"
)

type human struct {
	Sex int
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
		human: human{Sex: 1},
		//在teacher和student中使用了另一种匿名字段，
		//系统默认了human就是该字段的名称,虽然human也是其中的嵌套结构
	}
	b := student{Name: "Holo", Age: 18, human: human{Sex: 0}} //这样写好点
	fmt.Print(a, b)

	a.Name = "Liu"
	a.Age = 29
	a.human.Sex = 100

	fmt.Print(a, b)

	a.Sex = 10 //默认把嵌入的结构的属性都给个被嵌入的teacher和student。可以直接使用。
	//如果结构名相同，优先原结构，不考虑嵌套结构。
	fmt.Print(a, b)

}
