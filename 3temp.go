package main

import (
	"fmt"
	"math"
)

type (
	文本 string //将“文本”代替为string类型
)

func main() {
	var a int32
	var b float32
	var c string
	var d [2]int  //方扣号里表示数组容量，如果不指定容量就成了一个切片
	var e [3]byte //byte相当于unit8
	var f [4]rune //rune相当于int32
	var g [5]bool
	var h 文本
	h = "中文类型名" //在正式开发中尽量不要使用中文类型名

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(math.MaxInt32) //math包用于检查是否超出范围
	fmt.Println(h)
}
