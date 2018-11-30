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

func (c PhoneConnecter) Name() int { //这个结构参数是PhoneConnecter类型的，调用结构的名字叫做Name，无参数，返回数据是string
	fmt.Println("?", c.name)
	return 1 //return要放在后面。

}

func (c PhoneConnecter) Connect() { //这个结构参数是PhoneConnecter类型的，调用结构的名字叫做Connecter，返回就是打印的东西。
	fmt.Println("Connect:", c.name) //这里的c.name是上面来的，怎么能直接被使用呢？？？？
}

func main() {
	a := PhoneConnecter{"PhoneConnecter"} //相比之前的例子，简化了。因为USB这个接口里有关于PhoneConnecter这个结构类型的方法，所以a就可以调用接口里的方法。
	a.Connect()
	Disconnect(a)
}

func Disconnect(usb USB) { //传进来实现了USB接口的变量
	fmt.Println("Disconnected.") //如果能打印说明是实现了USB接口。
}

//-------------------
//视频老师讲的压根没调用Name（）string方法。
