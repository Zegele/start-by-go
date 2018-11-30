package main

import (
	"fmt"
)

func main() {
	var a int
	a = 1
	fmt.Println(a)

	var b int = 2 //感觉这个是最标准的
	fmt.Println(b)

	var c = 3 //类型推断，你确定系统不会推断错就可以使用。
	fmt.Println(c)

	d := "d"
	fmt.Println(d) //最简的声明和赋值，也是由系统推断类型的
}
