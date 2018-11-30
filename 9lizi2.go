package main

import (
	"fmt"
)

//传数值
func main() {
	a, b, c, e := 1, 2, 3, 4
	A(a, b, c)
	B(a, b, c, e)
	fmt.Println(a, b, c) //为什么是地址？
}

func A(b ...int) { //传进来的是slice,这个怎么理解？b是个slice可以这样理解？
	fmt.Println(1) //slice是引用类型，
	fmt.Println(b) //这里的b代表了A的所有参数？？？？
	fmt.Println(0)
}
func B(a, b, c, d int) { //与上面的不同是，这里我都写了具体的a,b,c,d。这里的b只是一个值
	fmt.Println(1) //slice是引用类型，
	fmt.Println(d) //这里的d,对应的是第四个参数，也就是在这里它对应的是e的值。
	fmt.Println(0)
}
