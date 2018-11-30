package main

import (
	"fmt"
	"reflect"
)

type ShiZhong struct {
	Name  string
	Count int
	KaiGuan
}

type KaiGuan struct {
	kaiguan int
}

func (u ShiZhong) DaoShu(tingzhi int) {
	fmt.Println(ShiZhong.Count)//不能打印出Count的值？就要用reflect?
	//i := ShiZhong.Count
	//fmt.Println(i)
	i:=100
	for i > tingzhi; i--{//复习了for循环
		fmt.Println(i)
}
}

func main() {
	Plan1 := ShiZhong{Name: "miqi", Count: 100, KaiGuan: KaiGuan{kaiguan: 1}}//复习了嵌入结构
	//注意！1.逗号，2.如果有嵌入结构，其他的名字也要写全。
	//写全，不能 ShiZhong{ "miqi", 100, KaiGuan: KaiGuan{kaiguan: 1}}这样，也就是不能省略Name，Count等
	fmt.Println(Plan1)
	Plan1.DaoShu(1)
	Plan1.Name="qimi"
	fmt.Println(Plan1)
}

func Set (o interface{}){
	v := reflect.ValueOf(o)
}

//来个倒数的方法。
//调用一个方法，将结构的初始值，改变为倒数截止的数字。