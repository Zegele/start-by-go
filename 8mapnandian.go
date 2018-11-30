package main

import (
	"fmt"
)

func main() {
	var m map[int]map[int]string     //value也是一个map多返回值
	m = make(map[int]map[int]string) //初始化第一层map,每一级map都要进行初始化。
	m[1] = make(map[int]string)      //初始化第二层map，同时指定是给key1的value初始化。
	m[1][1] = "OK"
	a := m[1][1]
	fmt.Println(a)
	fmt.Println(m)

	//多返回值
	b, c := m[2][1]   //第一个返回value，第二个当没有值时，返回false;有值，返回ture.
	fmt.Println(b, c) //打印看看m[2][1]是不是存在,
	if !c {           //如果c不存在
		m[2] = make(map[int]string)
	}
	m[2][1] = "GOOD GIRL!"
	b = m[2][1]
	fmt.Println(a, b, c)

	b, c = m[2][1] //多返回值,现在b对应的m[2][1]有值了，所以会返回ture
	fmt.Println(a, b, c)

}
