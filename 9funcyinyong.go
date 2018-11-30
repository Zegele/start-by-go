package main

import (
	"fmt"
)

func main() {
	a := A //这个A就是下面的函数A
	a()    //所以这就可以视函数为类型
}

func A() {
	fmt.Println("func A")
}
