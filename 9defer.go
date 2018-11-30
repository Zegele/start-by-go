package main

import (
	"fmt"
)

func main() {
	fmt.Println("a")
	defer fmt.Println("b")
	fmt.Println("c")
	defer fmt.Println("d")
	fmt.Println("e")
}

//acedb这个结果让人惊讶。。。
//当把其他打印结束后，才开始对defer的打印，切是倒序的
