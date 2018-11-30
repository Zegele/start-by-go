package main

import (
	"fmt"
)

func main() {

}

func A(a int, b string) (int, string) {
	//第一小括号里写，参数列表,如a int, b string，a、b是参数名称，int\string是参数类型。
	//第二个小括号是返回值,如果只有一个可以直接写不写该小括号如：func A(a int, b string)int

}

func B(a, b, c int) (e, f, g int) { //多个值都是int的写法
	//注意e,f,g已经被命名，所以不需要e :=1 ,应该是e=1

}

func C(a ...int) { //a ...int不定长变参，不指道a后还有多少，但知道它们都是int。
	//不定长变参只能放在最后。如(b string, a ...int)
	//a在接受一系列参数的时候变成了slice
	//不定长变参接收到的是指定类型的slice,那么和真正传递slice有什么区别？
	//实际上是进行的是值的拷贝，再形成slice,再被调用。
	//如果本身传递进来的是slice,是指针地址的拷贝，如果改变，都会改变原slice
}
