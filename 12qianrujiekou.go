//嵌入接口:接口里面套接口
package main

import (
	"fmt"
)

type USB interface { //例子：设置USB的接口
	//Name() string
	QIANRU //直接使用QIANRU接口
}
type QIANRU interface { //定义了QIANRU接口有个Connect()方法
	Connect()
}
type PhoneConnecter struct { //创建了一个名叫PhoneConnect的结构。
	name string
}

/*func (c PhoneConnecter) Name() string { //这个结构参数是PhoneConnecter类型的，调用结构的名字叫做Name，无参数，返回数据是string
	fmt.Println("?", c.name)
	return c.name //return要放在后面。

}

*/

func (c PhoneConnecter) Connect() { //这个结构参数是PhoneConnecter类型的，调用结构的名字叫做Connecter，返回就是打印的东西。
	fmt.Println("Connect:", c.name) //这里的c.name是上面来的，怎么能直接被使用呢？？？？
}

func main() {
	a := PhoneConnecter{"PhoneConnecter"} //相比之前的例子，简化了。但我们怎么知道实现这个接口呢？
	a.Connect()
	Disconnect(a)
}

func Disconnect(usb USB) { //传进来实现了USB接口的变量
	fmt.Println("Disconnected.") //如果能打印说明是实现了USB接口。
}

//-------------------
//视频老师讲的压根没调用Name（）string方法。
