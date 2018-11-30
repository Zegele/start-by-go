//通过反射调用方法
//感觉就就动态的调用方法

package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello(name string) {
	fmt.Println("Hello", name, ", my name is ", u.Name)
}

func main() {
	//-----------常规调用Hello方法
	u := User{1, "OK", 12}
	u.Hello("joe")

	//-----------使用反射调用Hello方法
	v := reflect.ValueOf(u)       //v反射了u的值
	mv := v.MethodByName("Hello") //取名字为Hello的方法。
	//v可能相应的有很多方法，你怎么知道有个方法名字叫做Hello呢？参照xiugai3
	//更严谨的是需要验证，这里不做严谨的演示。
	//只做调用方法的演示。

	args := []reflect.Value{reflect.ValueOf("joe")} //建立一个参数，参数要求必须传进去是slice，slice的元素是reflect.Value
	//slice的元素是reflect.Value
	mv.Call(args) //调用方法Call，把参数传进去。
}
