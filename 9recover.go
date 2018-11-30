package main

import (
	"fmt"
)

func main() {
	A()
	B() //让panic停止后，继续运行。
	C()
}

func A() {
	fmt.Println("func a")
}

func B() {
	defer func() {
		if err := recover(); err != nil { //检查panic是否存在
			//调用recover这个函数,会返回一个信息;
			//如果返回的信息不是nil,说明程序引发了panic,如果返回的是nil,说明什么都不需要做
			//再次强调，这个方法是检查程序是否引发了panic，且改变程序结果的操作。
			fmt.Println("Recover in B")
		}
	}()
	panic("Panic func b") //执行到这里，值应该就成了nil
	//defer(recover)要放在panic之前，因为程序遇到panic都不会进行下去了。

}

func C() {
	fmt.Println("func c")
}
func D() { //这里多一个没用的函数没关系？
	fmt.Println("func d")
}
