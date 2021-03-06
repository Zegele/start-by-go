//reflect反射
//通过ValueOf和TypeOf从 接口 中获得对象信息。

//演示了对一个结构的反射（字段，字段属性，方法，方法属性，类型，及类型名称等）
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

func (u User) NIUBI() {
	fmt.Println("NIUBIHUAILE.")
}

func (u User) Hello() { //创建一个方法
	fmt.Println("Hello world.")
}

func main() {
	u := User{1, "OK", 12}
	info(u) //传递的是值拷贝。
	u.Hello()
}

func info(o interface{}) { //创建一个方法，
	//用于传进一个空接口，然后输出这个传入结构的具体信息。

	//提取类型-------------------
	t := reflect.TypeOf(o)            //获得传入的接口的类型。
	fmt.Println("Type---:", t.Name()) //打印接口类型？
	//是TypeOf和ValueOf的输出值的Name，这个应该是有一套的（如Name，Type等）。
	fmt.Println(t) //t是mian.User。意思是主包的User类型？？

	//提取字段值-------------------
	v := reflect.ValueOf(o) //取字段（值）
	fmt.Println(v)          //打印取出的字段值，
	fmt.Println("Fields:")

	fmt.Println(t.NumField()) //打印t包含的个数

	for i := 0; i < t.NumField(); i++ { //t.NumField()
		f := t.Field(i)                                   //取t下包含的字段，如Id\Name\Age
		fmt.Println(f)                                    //打印字段
		val := v.Field(i).Interface()                     //取字段的值
		fmt.Println(val)                                  //取字段的值                     //Interface()方法是什么？？？
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val) //这个Name和Type已经不是结构内的Name了

		va := v.Field(i) //不需要Interface()方法不是也可以么？
		fmt.Println(va)

		fmt.Println("%6s: %v = %v", f.Name, f.Type, va)
		//Println是打印后面的内容
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, va)
		//Printf是按照这个格式打印
		fmt.Printf("%10s: %v = %v\n", f.Name, f.Type, va)
		//10代表后移位数
	}

	//提取方法-------------------
	for i := 0; i < t.NumMethod(); i++ { //提取t所拥有的方法。t.NumMethod()是方法个数。
		m := t.Method(i) //m是关于方法的详细描述，包含方法的Name，Type等信息。所以m.Name是打印具体的东西。
		fmt.Println(m)
		fmt.Printf("%6s：%v\n", m.Name, m.Type) //打印Name和方法签名[func(main.User)]
		//按字母排序
	}

}
