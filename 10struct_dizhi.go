package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	a := &person{
		Name: "Joson",
		Age:  19,
	} //在初始化的时候，对结构的初始化，习惯性使用取地址符号。
	//因为可以对原值进行操作。
	a.Name = "ok" //可以直接该原值，因为取的是地址。
	fmt.Println(a)
	A(a)
	fmt.Println(a)
}

func A(per *person) {
	per.Age = 12
	fmt.Println("A", per)
}
