package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 65
	b := string(a) //65在字符串里代表A
	fmt.Println(b)

	c := strconv.Itoa(a) //这个方法可以将65 转化为字符串的65
	fmt.Println(c)

	d, _ := strconv.Atoi(c) //将c字符串的65，转化成整型的65.
	//但是这个d后面的 逗号 和 杠 是什么意思？
	fmt.Println(d)
}
