package main

import (
	"fmt"
)

const (
	a = "A"
	b
	c = iota
	d
)

const (
	e = iota
)

const (
	FUCK = "f"
	YOU  = iota //一般命名全大写。
	//如果不想被包的外部使用，可以用 _YOU 或 cYOU(前加_或加小写字母)
	en
	AND = "and"
	you = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(FUCK)
	fmt.Println(YOU)
	fmt.Println(en)
	fmt.Println(AND)
	fmt.Println(you)
}
