package main

import (
	"fmt"
)

func main() {
	A()
	B() //程序运行到这里，panic停止了。
	C()
}

func A() {
	fmt.Println("func a")
}

func B() {
	panic("Panic func b")
}

func C() {
	fmt.Println("func c")
}
func D() { //这里多一个没用的函数没关系？
	fmt.Println("func d")
}
