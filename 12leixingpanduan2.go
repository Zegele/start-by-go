//类型判断
//ok pass 模式
package main

import (
	"fmt"
)

type empty interface { //空接口。。。可包含，容纳任何方法

}

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

func Disconnect(usb interface{}) { //传进空接口。
	switch v := usb.(type) {
	case PhoneConnecter: //如果是PhoneConnecter类型，那么……
		fmt.Println("Disconnected", v.name)
	default: //否则……
		fmt.Println("Unknown decive.")
	}
}
