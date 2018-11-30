package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{7, 8, 9}
	fmt.Println(s1, s2)
	copy(s1, s2)
	fmt.Println(s1, s2) //将s2 copy 到s1,将后拷贝到前

	s3 := []int{1, 2, 3, 4, 5, 6}
	s4 := []int{7, 8, 9}
	fmt.Println(s3, s4)
	copy(s4, s3) //盯住短的
	fmt.Println(s3, s4)

	copy(s3[2:4], s4[:2]) //copy到指定位置
	fmt.Println(s3, s4)
	s5 := s4[:] //取s4的从头到尾,s4完全相同
	fmt.Println(s5)
	fmt.Printf("%p\n", s4)
	fmt.Printf("%p", s5)
}
