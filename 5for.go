package main

import (
	"fmt"
)

func main() {
	a := 1 //for 第一种
	for {
		a++
		if a > 3 {
			break
		} //if语句块结束
		fmt.Println(a)
	} //for语句块结束
	fmt.Println("Over a")

	b := 2       //for 第二种，直接自己带条件表达式
	for b <= 3 { //b<=3（当b小于等于3，则执行） 就是自带的条件表达式，
		b++
		fmt.Println(b)
	}
	fmt.Print("Over B")

	c := 1                   //for 第三种 三个语句：1.如i=0的计数器初始化，2.i的条件语句，3.i的步进表达式
	for i := 0; i < 3; i++ { //计数器，条件表达式，步进表达式
		c++
		fmt.Println(c) //为什么这个打印结果会跟在之前的打印结果后？
	}
	fmt.Print("Over c")

	d := "Niubility"
	e := 10
	for i := 0; i < len(d); i++ { //最好不要在这里出现函数，因为每次都要进行运算，费时，费资源。
		e++
		fmt.Println(e)
	}
	fmt.Println("Over e")

	m := 10
	l := len(d) //将函数放在for外面，下面的 i < l,执行速度比上面的i<len(d)速度快。
	for i := 0; i < l; i++ {
		m++
		fmt.Println(m)
	}
	fmt.Println("Over m")
} //main包结束

/*
	for{

	}
这是无限循环，里面需要条件，否则程序卡死。
*/
