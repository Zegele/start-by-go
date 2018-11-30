package main

import (
	"fmt"
)

type TZ struct {
	Num int
}

func main() {
	a := TZ{Num: 0}
	fmt.Println(a)
	a.Print()
	fmt.Println(a)
}

func (a *TZ) Print() {
	for a.Num = 0; a.Num < 100; a.Num++ {
		//	fmt.Println(a.Num)
	}
	fmt.Println(a)
}
