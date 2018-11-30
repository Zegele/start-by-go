//给底层类型增加方法
package main

import (
	"fmt"
)

type TZ int //强制类型转换 因为底层的int类型无法进行绑定。因为不在一个包中是无法绑定的。
//把TZ叫做自定义类型。

func main() {
	var a TZ //不是结构，应该这样定义。 不能a := TZ 或 a := TZ{}后者我没想通为啥要这么写，我知道他写错了，但他脑子里在想什么
	a.Print()
	fmt.Println(a)
}

func (a *TZ) Print() { //给底层类型，附加了方法。附带方法有不同需求。
	var c TZ //或 var c TZ = 1
	c = 1
	fmt.Println(c)
	fmt.Println("tz")
}
这个receiver就是func的第一个参数，而且是强制性参数，方法和函数之间互通