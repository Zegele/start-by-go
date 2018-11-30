package main

import (
	"fmt"
)

const (
	B  float64 = 1 << (iota * 10) //浮点型的好处是可以科学记数法，整型可能会溢出
	KB                            //对哦，iota可以带入表达式内部哦。。。
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	//思路fmt.Println(1<<(n*10)) 变量只有n,然后返回上去在常量里设置。
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(MB)
	fmt.Println(ZB)
	fmt.Println(YB)

}
