package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	a := person{
		Name: "Joson",
		Age:  19,
	}
	b := person{Name: "Joson", Age: 19} //这样写 逗号可以省略？
	fmt.Println(b == a)

	/*
		c := person1{Name: "Joson", Age: 19}
		fmt.Println(b == c)
	*/
	//名称不同就是不同的类型，无法比较

	d := person{Name: "Joson", Age: 20} //名字相同，值不同，则不相等。
	fmt.Println(b == d)
}
