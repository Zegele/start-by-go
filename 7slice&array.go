package main

import (
	"fmt"
)

func main() {
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'g', 'k'}
	fmt.Println(a) //这是asca码
	fmt.Println(string(a))

	sa := a[2:5]
	fmt.Println(string(sa))

	sb := a[3:5]
	fmt.Println(string(sb))
}
