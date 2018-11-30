package main

import (
	"fmt"
)

type person struct {
	Name    string
	Age     int
	Contact struct { //匿名结构中，嵌套了匿名结构
		Phone         int
		City, Address string //简写，不用重复打回车，打string类型。
	}
}

func main() {
	a := person{ //值的初始化
		Name: "Liu",
		Age:  29, //这个逗号也不能少
	}
	a.Contact.Phone = 18113977650 //嵌套的匿名结构的值的初始化
	a.Contact.City = "乐山"
	a.Contact.Address = "千佛岩"
	fmt.Println(a)
}
