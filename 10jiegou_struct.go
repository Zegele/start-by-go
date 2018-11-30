package main

import (
	"fmt"
)

//type test struct{} //定义了空结构,test 和person是自己命名的。
type person struct {
	Name string
	Age  int
}

/*func main() {
	a := person{}
	fmt.Println(a)
}
默认值是零值*/

func main() {
	a := person{} //如果a:= &person{}取person的指针。
	a.Name = "Joson"
	a.Age = 19
	fmt.Println(a)

	//简便写法
	b := person{
		Name: "Liu", //字面值的初始化
		Age:  29,    //注意  ， 逗号 别忘
	}
	fmt.Println(b)

	A(a)

	fmt.Println(a)

	B(&a) //取a的地址。

	fmt.Println(a) //把原地址的值也改变了。
}

func A(per person) { //函数A，参数per(这个per可以任意)，是person类型
	per.Age = 13
	fmt.Println("A", per) //打印per 说明是值的拷贝

}

func B(bb *person) { //person的指针
	bb.Age = 14
	bb.Name = "jjjjj"
	fmt.Println("A", bb) //打印bb说明值改变了。

}
