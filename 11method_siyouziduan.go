//方法的访问权限
package main

import (
	"fmt"
)

type A struct {
	name string //只要在同一个包内都可以被  方法method  访问。
}

func main() {
	a := A{}
	fmt.Println(a)
	a.Print()
	fmt.Println(a)
	fmt.Println(a.name)
}

func (a *A) Print() {
	a.name = "123"
	fmt.Println(a.name)
}
