package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 3, 6)
	fmt.Printf("%p\n", s1) //%p打印s1的内存地址? \n换行符号后换行
	fmt.Println(len(s1), cap(s1))

	s1 = append(s1, 1, 2, 3)      //给s1加了1，2，3，三个元素
	fmt.Printf("%v %p\n", s1, s1) //%v打印值？
	fmt.Println(len(s1), cap(s1))

	s1 = append(s1, 1, 2, 3) //给s1又又又加了1，2，3，三个元素，超出了
	fmt.Printf("%v %p\n", s1, s1)
	fmt.Println(len(s1), cap(s1))
}
