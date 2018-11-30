//接口
package main

import (
	"fmt"
)

type USB interface { //例子：设置USB的接口
	Name() string //为什么加string？
	Connect()
}

type PhoneConnecter struct { //创建了一个名叫PhoneConnect的结构。
	name string
}

func (c PhoneConnecter) Name() string { //这个结构参数是PhoneConnecter类型的，调用结构的名字叫做Name，无参数，返回数据是string
	fmt.Println("?", c.name)
	return c.name //return要放在后面。

}

func (c PhoneConnecter) Connect() { //这个结构参数是PhoneConnecter类型的，调用结构的名字叫做Connecter，返回就是打印的东西。
	fmt.Println("Connect:", c.name) //这里的c.name是上面来的，怎么能直接被使用呢？？？？
}

func main() {
	b := PhoneConnecter{}
	b.name = "Liu"
	fmt.Println(b)

	var a USB                                  //定义a是USB类型，a是一个接口
	a = PhoneConnecter{name: "PhoneConnecter"} //将结构赋值给a，并给了一个初始值。
	fmt.Println(a)

	a.Connect() //调用a接口的Connect方法
	a.Name()    //调用a接口的Name方法

}
