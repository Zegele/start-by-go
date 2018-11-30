package main

import (
	"fmt"
)

func main() {
	f := closure(10) //声明函数f，并调用闭包的外部函数，这个闭包函数会返回一个匿名函数。
	fmt.Println(f(1))
	fmt.Println(f(2))
}

func closure(x int) func(int) int { //闭包也是一个函数
	//1.closure 相当于函数名
	//2.（x int) x相当于函数参数，int是参数类型
	//3.func(int) int 相当于返回值，而这个返回值恰好是个函数
	fmt.Printf("%p\n", &x)   //在闭包中打印地址
	return func(y int) int { //闭包的作用是返回 一个匿名函数 ！！！
		fmt.Printf("%p\n", &x) //在返回函数中打印x地址
		//x的地址没变，说明用的是同一个x的值
		return x + y //x来自于外层函数，y来自内层匿名函数。
	}
}
