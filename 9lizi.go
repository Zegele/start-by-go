package main

import (
	"fmt"
)

//传数值
func main() {
	A(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(A) //为什么是地址？
}

func A(a ...int) { //传进来的是slice,这个怎么理解？
	fmt.Println(a) //slice是引用类型，
}
