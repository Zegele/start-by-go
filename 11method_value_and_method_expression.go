//method value 和 method expression 的对比，有相同的效果。

package main

import (
	"fmt"
)

type TZ int

func main() {
	var a TZ
	a.Print()       //方法和函数之间互通1 method value:调用已申明的方法
	(*TZ).Print(&a) //方法和函数之间互通2method expression：直接通过类型（而不是某个类型的变量进行调用。上面就是利用变量a调用Print方法）
	//类型.方法（需要自己传进来一个receiver）将receiver作为第一个参数传递给方法。
	//因为这里传进来的是指针，所以要取地址。
	//这两种形式都调用了Print方法。
}

func (a *TZ) Print() {
	//这个receiver（a *TZ）就是func的第一个参数，而且是强制性参数，方法和函数之间互通
	var c TZ
	c = 1
	fmt.Println(c)
	fmt.Println("tz")
}
