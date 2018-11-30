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
	fmt.Println(v)                                     //这个是指针。。这个描述我不知道对不对
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
	}
	fmt.Println(v) //这个是重新赋过值的

	f := v.FieldByName("Name") //这个Name在不在User里？？？就需要判断
	if !f.IsValid() {          //判断是不是真的找到了这个字段。如果不是IsValid的则打印BAD
		//如果没有找到，会返回一个空的reflect.value
		fmt.Println("BAD")
		return
	}

	if f.Kind() == reflect.String {
		f.SetString("NIUBI")
	}

	z := v.FieldByName("Age") //这个Age在不在User里？？？就需要判断
	if !z.IsValid() {         //判断是不是真的找到了这个字段。如果不是IsValid的则打印BAD
		//如果没有找到，会返回一个空的reflect.value
		fmt.Println("BADBAD")
		return
	}
	if z.Kind() == reflect.Int {
		z.SetInt(100)
	}

	y := v.FieldByName("AgeAGE") //这个AgeAGE在不在User里？？？就需要判断
	if !y.IsValid() {            //判断是不是真的找到了这个字段。如果不是IsValid的则打印BAD
		//如果没有找到，会返回一个空的reflect.value
		fmt.Println("GOOD")
		fmt.Println(v)
		fmt.Println(f)
		fmt.Println(z)
		fmt.Println(y)
		return //return之后不执行了？了？了？
	}

	if y.Kind() == reflect.Int {
		y.SetInt(100)
	}

}
