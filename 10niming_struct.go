package main

import (
	"fmt"
)

func main() {
	a := struct {
		Name string
		Age  int
	}{
		Name: "Jim",
		Age:  19,
	}
	fmt.Println(a)

	b := &struct { //取地址 &
		Name string
		Age  int
	}{ //紧跟着就可以赋值，字面值的初始化。
		Name: "Tom",
		Age:  19,
	}
	fmt.Println(b)
	b.Name = "Liu"
	b.Age = 30
	fmt.Println(b)
}
