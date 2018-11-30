package main

import (
	"fmt"
)

type A struct {
	Num int
}

type B struct {
	Name string
}

func main() {
	a := A{}  //声明a是结构A
	a.Print() //调用方法记得加（）小括号
	fmt.Println(a)

	b := B{}  //声明b是结构B
	b.Print() //虽然表面看a 和b 的方法名字一样，但其实不一样的。
	//因为a和b各自绑定的是不同的结构类型（或其他类型），所以不一样
	fmt.Println(b)

	//----------传递指针 的演示
	c := A{}
	c.Prin() //*c.Prin() 这个*号不用加。
	fmt.Println(c)

	a.Pri() //不用 *a.Pri() 这个*号不用加。系统会自动根据你写的程序判断，这个值是指针值，还是拷贝值。
	fmt.Println(a.Num)

}

//方法  有个接收者（receiver)的概念

func (a A) Print() { //为a写了一个 结构（struct)  的方法 Print方法。
	//1.(a A)括号里就是接收者，a是局部变量，A 就代表类型，这里类型就是结构(struct)A
	//2.取方法（method）的名称 如Print
	//根据接收者的类型判断是哪个结构的方法。
	//
	fmt.Println("1000") //并没有对struct进行操作。仅仅是个打印而以
	a.Num = 100         //对struct进行操作,但并没有对原结构造成影响。所以这是对值的拷贝。
	fmt.Println(a)

}

/*
func (a A) Print() {}
从函数角度看，Print是函数名称，（）是函数参数。怎么解释？？？

*/

func (b B) Print() {
	fmt.Println(1000) //并没有对struct进行操作。仅仅是个打印而以
	b.Name = "NIUBI"  //对struct进行操作,但并没有对原结构造成影响。所以这是对值的拷贝。
	fmt.Println(b)
}

//---------------------------传递指针 的演示
//（a A）也是参数，是参数就涉及是值传递（对原始数据没影响），还是指针传递（会修改原始数据）。如下
func (c *A) Prin() { //指针传递
	fmt.Println("1000")
	c.Num = 100
	fmt.Println(c)
}
func (a *A) Pri() {
	fmt.Println("1000")
	a.Num = 100
	fmt.Println(a)
}

//注意：一个变量只能绑定一个类型的方法，这个方法的名字只能被使用一次。
/*如
a A Print
b A Print
这样就不行；

但
a A Print
b B Print
这样就行；

或
a A Print
b A Prin
这样就行
*/
