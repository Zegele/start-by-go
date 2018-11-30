package main

import (
	"fmt"
)

type TZ int

func main() {
	var a TZ
	fmt.Println(a)
	a.Print()       //方法1
	(*TZ).Print(&a) //方法2：类型后有个  .点  切记！！！
}

func (a *TZ) Print() { //这TM跟a几乎没啥关系。就是创造了个方法而以。a相当于一把钥匙，方法里是一个世界，世界和钥匙的关系就是，钥匙打开世界的那一刹那。
	var i TZ = 0
	for i = 0; i < 100; i++ {
	}
	fmt.Println(i)
}
