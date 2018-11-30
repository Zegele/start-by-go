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
	Kaiguan int //注意kaiguan小写不能被外部引用。首字母大写
}

func (u ShiZhong) DaoShu(TingZhi int) {
	fmt.Println(u.Count) //注意是u.Count,不是ShiZhong.Count!!!//可省略
	i := u.Count
	switch { //复习了switch
	case i > TingZhi:
		for i := u.Count; i > TingZhi; i-- { //i := u.Count 这个不能放在for循环外面。
			fmt.Println(i)
		}
	case i == TingZhi:
		fmt.Println(i)
	default:
		for i := u.Count; i < TingZhi; i++ { //i := u.Count 这个不能放在for循环外面。
			fmt.Println(i)
		}
	}
}

func main() {
	Plan1 := ShiZhong{Name: "miqi", Count: 10, KaiGuan: KaiGuan{Kaiguan: 0}} //复习了嵌入结构
	Plan1.DaoShu(20)                                                         //98是停止数字
	//Plan1.Name = "qimi"//复习直接修改值//可省略
	//这样就直接修改了原结构的值，但现在这里不用这种方法。//可省略
	//fmt.Println(Plan1)//可省略

	Plan2 := Plan1.KaiGuan //Plan1.KaiGuan是结构类型；Plan1.KaiGuan.kaiguan是int类型
	fmt.Println(Plan2)     //{0}
	Set(&Plan2)
	fmt.Println(Plan2) //{11} 经过Set(&Plan2)已经修改。
	fmt.Println(Plan1) //注意并没有改变Plan1的值。

	Plan3 := ShiZhong{Name: "miqi", Count: 100, KaiGuan: KaiGuan{Kaiguan: 1}} //为了好识别加上的，可省略
	fmt.Println(Plan3)
	Set(&Plan3)
	fmt.Println(Plan3)
}

//--------------取值
/*func Quzhi(q interface{}) {
	t := reflect.TypeOf(q)
	fmt.Println("Type:", t.Name())
	for i := 0; i < t.NumField(); i++ {
		z := t.Field(i)

	}

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("传入类型错误，转化失败")
		return
	}
}
*/
//--------------改值
func Set(q interface{}) { //这里必须是接口类型么？？？
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)        //因为这里是指针，所以。。。
	fmt.Println("Type:", t.Name()) //可省略
	for i := 0; i < t.NumField(); i++ {
		l := t.Field(i)
		fmt.Println(l.Name)
	}

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() { //如果v.Kind()==reflect.Ptr成立 或 !v.Elem().CanSet()成立，则进行。。。
		fmt.Println("不可更改")
		return
	} else {
		v = v.Elem() //所以这里重新赋值了一下//括号里不填写东西吗？？？
	}
	fmt.Println(v)

	//注意嵌套结构内的名字能直接用reflect(反射)修改值-----------------
	f := v.FieldByName("Kaiguan") //如果找不到这个名字返回的是空值。
	fmt.Println(f)
	if !f.IsValid() {
		fmt.Println("不可设置")
		return
	}
	if f.Kind() == reflect.Int {
		f.SetInt(1001)
	}
}

//来个倒数的方法。
//调用一个方法，将结构的初始值，改变为倒数截止的数字。
//用反射方法把结构的初始值全部修改掉。
