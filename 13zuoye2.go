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

func (u ShiZhong) DaoShu(TingZhi int) {
	fmt.Println(u.Count) //注意是u.Count,不是ShiZhong.Count!!!//可省略

	for i := u.Count; i > TingZhi; i-- { //i := u.Count 这个不能放在for循环外面。
		fmt.Println(i)
	}
	geshu1 := u.Count - TingZhi
	fmt.Println(geshu1)
}

func main() {
	Plan1 := ShiZhong{Name: "miqi", Count: 100, KaiGuan: KaiGuan{kaiguan: 0}} //复习了嵌入结构
	//注意！1.逗号，2.如果有嵌入结构，其他的名字也要写全。
	//写全，不能 ShiZhong{ "miqi", 100, KaiGuan: KaiGuan{kaiguan: 1}}这样，也就是不能省略Name，Count等
	fmt.Println(Plan1)
	fmt.Println(Plan1.KaiGuan)         //可省略
	fmt.Println(Plan1.KaiGuan.kaiguan) //可省略
	Plan1.DaoShu(98)                   //98是停止数字
	//Plan1.Name = "qimi"//复习直接修改值//可省略
	//这样就直接修改了原结构的值，但现在这里不用这种方法。//可省略
	//fmt.Println(Plan1)//可省略

	Plan2 := Plan1.KaiGuan //Plan1.KaiGuan是结构类型；Plan1.KaiGuan.kaiguan是int类型
	fmt.Println(Plan2)

	Plan3 := ShiZhong{Name: "miqi", Count: 100, KaiGuan: KaiGuan{kaiguan: 1}} //为了好识别加上的，可省略
	fmt.Println(Plan3)
	Set(&Plan3)
	fmt.Println(Plan3)
}

func Set(q interface{}) { //这里必须是接口类型么？？？
	v := reflect.ValueOf(q)                            //因为这里是指针，所以。。。
	fmt.Println(v)                                     //这里是指针的值
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() { //如果v.Kind()==reflect.Ptr成立 或 !v.Elem().CanSet()成立，则进行。。。
		fmt.Println("不可更改")
		return
	} else {
		v = v.Elem() //所以这里重新赋值了一下//括号里不填写东西吗？？？
	}
	fmt.Println(v)

	f := v.FieldByName("KaiGuan") //没找到这个名字返回的是空值。
	if !f.IsValid() {
		fmt.Println("不可设置")
		return
	}
	if f.Kind() == reflect.Int {
		f.SetInt(11)
	}

}

//来个倒数的方法。
//调用一个方法，将结构的初始值，改变为倒数截止的数字。
//用反射方法把结构的初始值全部修改掉。
