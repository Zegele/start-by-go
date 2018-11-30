package main

import (
	"fmt"
)

type person struct {
	string //匿名字段 Name string
	int
}

func main() {
	a := person{"Tom", 19} //只能按照上面的顺序初始化或输入。
	fmt.Println(a)
}
