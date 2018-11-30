package main

import (
	"fmt"
)

func main() {
	//第一种方法创建map
	var m map[int]string //[key]value
	m = map[int]string{} //这个操作是必须的吗？初始化
	fmt.Println(m)

	//第二种方法创建map
	/*var n map[int]string
	  n = make(map[int]string)
	  简写如下，两行缩写成一行*/
	var n map[int]string = make(map[int]string)
	//map[int]string这一串可以视为类型(type)
	fmt.Println(n)

	//第三种方法创建map，最简
	o := make(map[int]string)
	//map[int]string可以视为类型（type)
	fmt.Println(o)

	o[1] = "OK"    //给key为1的取值OK
	o[2] = "NIUBI" //给map o中1，存值“OK”
	fmt.Println(o)
	a := o[1] //a取map o中1的值。
	fmt.Println(a)

	delete(o, 2) //删除map o 中2键的值
	fmt.Println(o)
}
