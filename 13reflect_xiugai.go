//通过反射修改基本类型内容

package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 123
	fmt.Println(x)

	v := reflect.ValueOf(&x) //取出x的地址，传入指针。
	fmt.Println(v.Elem())    //v.Elem()是reflect value

	v.Elem().SetInt(999) //传入指针就能调用这个方法。

	fmt.Println(x) //然后打印出来的值就是改变的值。
}
