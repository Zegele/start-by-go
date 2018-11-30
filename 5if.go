package main

import (
	"fmt"
)

func main() {
	a := 10 //进入if,因为a重名，所以这个外部的a会被隐藏。
	if a := 1; 1 < 2 {
		fmt.Println(a) //输出if语句块内的a
		//if内的a只在if语句块内有效
		//0. if然后空格，直接条件表达式。不需要括号。
		//1.a:=1是初始化语句；
		//2.左大括号必须放在if同一行
		//3.；分号 隔开初始语句与条件语句
	}
	fmt.Println(a) //输出if外的a
}
