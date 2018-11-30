package main

import (
	"fmt"
)

//传数值
func main() {
	//原值没有变
	a, b, c, e := 1, 2, 3, 4
	A(a, b, c, e) //把abcd传递给函数A
	fmt.Println(a, b, c, e)

	//原址被改变了
	s1 := []int{10, 11, 12, 13}
	B(s1)           //把s1（slice）传递给函数A
	fmt.Println(s1) //打印了已经被改的s1
}

func A(s ...int) { //这不是个slice
	//值的拷贝
	s[0] = 30
	s[1] = 40
	fmt.Println(s) //返回的是一个slice,所以改变不了原来的abcd的值，因为abcd是数，不是slice类型。
}

func B(s []int) { //这是个slice，这里要求参数是个slice类型的，所以传进来的必须是slice
	//是对内存地址的拷贝，如果拿到了slice的地址，那就是对它本身在进行操作。
	s[0] = 30
	s[1] = 40
	fmt.Println(s) //已经把原来的s1改掉了。
	fmt.Println(0)
}

//是否改变原来的值，主要看是拷贝值，还是拷贝地址。
