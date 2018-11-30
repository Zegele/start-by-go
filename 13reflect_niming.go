//反射 匿名或嵌入字段
//--------------反射结构的信息。

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

type Manager struct {
	User  //嵌入字段，嵌入结构/User/反射会将匿名字段当做独立字段进行处理。 名字Name叫User，类型Type还叫
	title string
}

func main() {
	m := Manager{User: User{1, "OK", 12}, title: "123"}
	t := reflect.TypeOf(m) //取出类型

	fmt.Printf("%#v\n", t.Field(0)) //取出第一个字段
	fmt.Println(t.Field(0).Anonymous)//第1个是匿名的么？true是的，false不是
	fmt.Printf("%#v\n", t.Field(1)) //取出第二个字段
	//反射包使用序号的方式取出匿名字段

	//取出匿名或嵌入结构中的值---------------
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0})) //取Manager中的Use的ID
	//[]int{0,0})相对Manager来讲User是第0个，相对User来讲ID是第0个。这是索引
	//这就打印出ID，及有关信息。
	//相对于User来讲，ID，Name，Age都不是匿名字段。
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 2}))
	//这就打印出Age，及有关信息。
}
