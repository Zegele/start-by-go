package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	s1 := a[1:3]
	s2 := a[2:5]
	fmt.Println(s1, s2)
	fmt.Println(len(s1), cap(s1), len(s2), cap(s2))
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", s2)

	s2[0] = 9           //将s1中的第一个元素设为9，即3变9
	fmt.Println(s1, s2) //这时两个slice对应的数都会变。
	fmt.Println(len(s1), cap(s1), len(s2), cap(s2))
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", s2)

	s1 = append(s1, 1, 1, 1) //但当s2超出后，会分配新的内存地址,翻倍，就不会受s1的影响了。
	s2[0] = 8                //这时候改s2中的数不会影响s1了。
	fmt.Println(s1, s2)
	fmt.Println(len(s1), cap(s1), len(s2), cap(s2))
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p", s2)
}
