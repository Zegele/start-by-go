package main

import (
	"fmt"
)

func main() {
	a := func() { //匿名函数
		fmt.Println("func A")
	} //这一部分都是在定义a是这个函数。

	a() //程序执行到这里才开始打印"func c"
	//这个a()一定要注意，一定是a(),不能只是a
	fmt.Println(a) //这里应该是函数a的地址。
	/*a
	如果这里放a，会报错，因为没有使用a,函数就得有个函数的样子，如a()
	*/
}
