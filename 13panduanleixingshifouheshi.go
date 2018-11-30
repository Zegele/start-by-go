//判断是否传进去合适或希望或想要的类型？？

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

func (u User) NIUBI() {
	fmt.Println("NIUBIHUAILE.")
}

func (u User) Hello() {
	fmt.Println("Hello world.")
}

func main() {
	u := User{1, "OK", 12}
	info(u)
	info(&u) //取出u的地址，传进去一个指针。传进来的类型不对。指针类型不是

	u.Hello()
}

func info(o interface{}) {

	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name()) //转化失败是空的值。所以打印结果“Type：” 后是空的

	//判断传入的类型是否错误。
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("类型错误，转化失败")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s：%v\n", m.Name, m.Type)
	}

}
