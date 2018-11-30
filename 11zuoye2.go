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
	a.Print()       //方法1
	(*TZ).Print(&a) //方法2：类型后有个  .点  切记！！！
}

func (a *TZ) Print() {
	for a.Num = 0; a.Num < 100; a.Num++ {
		//	fmt.Println(a.Num)
	}
	fmt.Println(a)
}
