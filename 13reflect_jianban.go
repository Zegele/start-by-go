//reflect反射
//通过ValueOf和TypeOf从 接口 中获得对象信息。

package main

import (
	"fmt"
	"reflect" //引入反射包
)

type User struct { //创建一个结构
	Id   int
	Name string //后面的t.Name和f.Name不是这个Name
	Age  int
}

func (u User) Hello() { //创建一个方法
	fmt.Println("Hello world.")
}

func (u User) NIUBI() { //创建一个方法
	fmt.Println("NIUBIHUAILE.")
}

func main() {
	u := User{1, "OK", 12}
	info(u)
	u.Hello()
}

func info(o interface{}) { //创建一个方法，
	//用于传进一个空接口，然后输出这个传入结构的具体信息。
	t := reflect.TypeOf(o)         //获得传入的接口的类型。
	fmt.Println("Type:", t.Name()) //打印接口类型？
	//是TypeOf和ValueOf的输出值的Name，这个应该是有一套的（如Name，Type等）。

	v := reflect.ValueOf(o) //取字段（值）
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ { //t.NumField()
		f := t.Field(i)                                   //取t下包含的字段，如Id\Name\Age
		val := v.Field(i).Interface()                     //取字段的值
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val) //这个Name和Type已经不是结构内的Name了
	}

	for i := 0; i < t.NumMethod(); i++ { //提取t所拥有的方法。t.NumMethod()是方法个数。
		m := t.Method(i)
		fmt.Printf("%60s：%v\n", m.Name, m.Type)
	}
}
