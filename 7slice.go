package main

import (
	"fmt"
)

//创建slice
func main() {
	var s1 []int //创建了int的切片类型。
	fmt.Println(s1)

	s11 := [...]int{1, 2} //创建了int的切片类型。用[...]就必须跟字面值，如{1，2}
	fmt.Println(s11)

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(a)
	s2 := a[9] //在数组a中截取索引为九的元素
	fmt.Println(s2)

	s3 := a[5:10] //截取5到9的索引元素。注意不包含10.类似数学的[5,10)
	fmt.Println(s3)

	s4 := a[5:len(a)] //取到a的最后一位
	fmt.Println(s4)

	s5 := a[5:]
	fmt.Println(s5) //取到a的最后一位

	//s6 := a[0:5]//从a的第一位取值
	s6 := a[:5] //从a的第一位取值
	fmt.Println(s6)

	//正式的创建slice,可以指定slice完整的参数。
	s7 := make([]int, 3, 10)
	//1.make关键字；2.[]；3.类型，如int；4.指定元素数；5.元素容量，如分配10个内存地址，如果加超后，容量会翻倍。
	fmt.Println(s7)
	fmt.Println(len(s7), cap(s7))
}
