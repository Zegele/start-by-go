//通过反射修改结构类型内容

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

func main() {
	u := User{1, "OK", 12}
	fmt.Println(u)
	Set(&u) //取出U的地址，传入指针
	fmt.Println(u)
}

func Set(o interface{}) { //要被修改的对象，要求接口必须是可修改的，然后需要传指针给接口
	v := reflect.ValueOf(o)
	fmt.Println("123", v)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() { //判断取出的值是不是point interface类型。普通类型？？参照xiugai1看看
		//如果不是point interface操作是不会成功的。
		//CanSet（）方法会返回一个bull值，看传入的是不是可以被修改。
		fmt.Println("XXX")
		return
		//如果v.Kind() == reflect.Ptr，或 不是v.Elem().CanSet()则fmt.Println("XXX")，否则...
	} else { //如果以上两个都满足（！=并且是CanSet可修改的），则。。。。
		v = v.Elem() //为什么这个v可以被 v.Elem()赋值？
		//因为ValueOf(o)返回的值v是reflect.value,v.Elem()的值也是reflect.value，类型相同，则可赋值。
		//这里就取得了可赋值的实际的对象。
		fmt.Println(v)
	}

	if f := v.FieldByName("Name"); f.Kind() == reflect.String { //索引，之前是通过数字索引，这里通过"Name"索引
		//通过索引名称取出字段，v.FieldByName("Name")，取出Name的字段
		// Kind()方法判断取出来的值是否是reflect.String  f.Kind()==reflect.String
		f.SetString("BYEBYE") //如果是reflect.String则进行这个修改。
		//Kind()方法之前遇到过。
		//问题我们怎么知道“Name”字段在不在结构里呢？
	}
}
