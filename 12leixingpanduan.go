//类型判断
//ok pass 模式
package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connecter
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

//-----------------------------------------

type Connecter interface {
	Connect()
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connected:", pc.name)
}

//---------------------------------------------

func main() {
	a := PhoneConnecter{"PhoneConnecter"}
	a.Connect()
	fmt.Println(a)
	Disconnect(a)
}

func Disconnect(usb USB) { //
	if c, ok := usb.(PhoneConnecter); ok { //这句什么意思，完全不懂...c代表什么。。。
		//当ok判断完后，判断后的值就赋值给c,这才有了c.name
		fmt.Println("Disconnected:", c.name)
		return
	}
	fmt.Println("Unknown decive.")
}
