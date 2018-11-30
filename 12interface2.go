//接口
package main

import (
	"fmt"
)

type USB interface { //例子：设置USB的接口
	//Name() int //为什么加string？
	Connect()
}

type PhoneConnecter struct { //创建了一个名叫PhoneConnect的结构。
	name string
}

/*func (c PhoneConnecter) Name() int { //这个结构参数是PhoneConnecter类型的，调用结构的名字叫做Name，无参数，返回数据是string
	fmt.Println("?", c.name)
	return 1 //return要放在后面。

}*/

func (c PhoneConnecter) Connect() { //这个结构参数是PhoneConnecter类型的，调用结构的名字叫做Connecter，返回就是打印的东西。
	fmt.Println("Connect:", c.name) //这里的c.name是上面来的，怎么能直接被使用呢？？？？
}

func main() {
	b := PhoneConnecter{}
	b.name = "Liu"
	fmt.Println(b)

	var a USB                                  //定义a是USB类型，a是一个接口
	a = PhoneConnecter{name: "PhoneConnecter"} //将结构赋值给a，并给了一个初始值。
	//由于a已经定义为USB类型了，所以在赋值的时候最好直接给初始值。
	//
	fmt.Println(a)
	a = b
	a.Connect() //调用a接口的Connect方法

}

//-------------------
//视频老师讲的压根没调用Name（）string方法。
