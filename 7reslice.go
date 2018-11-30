package main

import (
	"fmt"
)

func main() {
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'g', 'k'}
	sa := a[2:5]
	fmt.Println(string(sa), len(sa), cap(sa))
	sb := sa[1:] //在slice a中取出slice b,按a数
	fmt.Println(string(sb), len(sb), cap(sb))
	//reslice 超出数组a,会报错，不会重新分配内存。

}
