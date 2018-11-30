package main

import (
	"fmt"
)

type A struct { //1-1
	B
	Name string
}

type B struct { //1-2
	Name string
}

//----------------------------------------

type C struct { //2-1
	D
}

type D struct { //2-2
	Name string
}

//-----------------------------------

/*
type X struct {//3-1
	Y
	Z//YZ因为同名（都是Name），所以不能识别。
}

type Y struct {//3-2
	Name string
}

type Z struct {//3-3
	Name string
}
*/

//---------------------------------
type X struct { //3-1
	Y
}

type Y struct { //3-2
	Z
	Name string
}

type Z struct { //3-3
	Name string
}

//-------------------------------------

func main() {
	a := A{Name: "A", B: B{Name: "B"}}
	fmt.Println(a)
	fmt.Println(a.Name, a.B.Name)

	c := C{D: D{Name: "D"}}
	fmt.Println(c)
	fmt.Println(c.Name, c.D.Name) //这里的c.Name就是c.D.Name，可以把D中的Name直接使用。可参照上一篇。

	y := Y{Z: Z{Name: "Z"}, Name: "Y"}
	fmt.Println(y)
	fmt.Println(y.Name, y.Z.Name)

	//--------------------------------------

	x := X{y} //因为X结构声明中是直接包着Y结构的，所以这里的大括号可以直接放个完整的结构。
	//我做了尝试只能定义一层结构，定义不了结构下一层的结构。
	fmt.Println(x)
	fmt.Println(x.Name, x.Y.Name)
}

//-----------------------------------
