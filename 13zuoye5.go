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
	//注意！1.逗号，2.如果有嵌入结构，其他的名字也要写全。
	//写全，不能 ShiZhong{ "miqi", 100, KaiGuan: KaiGuan{kaiguan: 1}}这样，也就是不能省略Name，Count等
	fmt.Println(Plan1)
	fmt.Println(Plan1.KaiGuan)         //可省略
	fmt.Println(Plan1.KaiGuan.Kaiguan) //可省略
	Plan1.DaoShu(20)                   //98是停止数字
	//Plan1.Name = "qimi"//复习直接修改值//可省略
	//这样就直接修改了原结构的值，但现在这里不用这种方法。//可省略
	//fmt.Println(Plan1)//可省略

	//利用反射调用DaoShu方法-----------------
	v := reflect.ValueOf(Plan1)
	mv := v.MethodByName("DaoShu")
	args := []reflect.Value{reflect.ValueOf(1)}
	mv.Call(args)

	Plan2 := Plan1.KaiGuan //Plan1.KaiGuan是结构类型；Plan1.KaiGuan.kaiguan是int类型
	fmt.Println(Plan2)     //{0}
	Set1(&Plan2)
	fmt.Println(Plan2) //{11} 经过Set(&Plan2)已经修改。
	fmt.Println(Plan1) //注意并没有改变Plan1的值。

	Plan3 := ShiZhong{Name: "miqi", Count: 100, KaiGuan: KaiGuan{Kaiguan: 1}} //为了好识别加上的，可省略
	fmt.Println(Plan3)
	Set2(Plan3)
	fmt.Println(Plan3)
	Set1(&Plan3)
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
func Set1(q interface{}) { //这里必须是接口类型么？？？
	v := reflect.ValueOf(q)                            //因为这里是指针，所以。。。
	fmt.Println(v)                                     //这里是指针的值
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() { //如果v.Kind()==reflect.Ptr成立 或 !v.Elem().CanSet()成立，则进行。。。
		fmt.Println("不可更改")
		return
	} else {
		v = v.Elem() //所以这里重新赋值了一下//括号里不填写东西吗？？？
	}
	fmt.Println(v)

	//注意嵌套结构内的名字能直接用reflect(反射)修改值-----------------
	f := v.FieldByName("Kaiguan") //如果找不到这个名字返回的是空值。
	fmt.Println(f)                //这里的f是0，直接是0//可省略
	if !f.IsValid() {
		fmt.Println("不可设置")
		return
	}
	if f.Kind() == reflect.Int {
		f.SetInt(1001)
	}

	u := v.FieldByName("Name") //如果找不到这个名字返回的是空值。
	fmt.Println(u)             //这里的f是0，直接是0//可省略
	if !u.IsValid() {
		fmt.Println("不可设置")
		return
	}
	if u.Kind() == reflect.String {
		u.SetString("已修改")
		fmt.Println(u)
	}

	c := v.FieldByName("Count") //如果找不到这个名字返回的是空值。
	fmt.Println(c)              //这里的f是0，直接是0//可省略
	if !c.IsValid() {
		fmt.Println("不可设置")
		return
	}
	if c.Kind() == reflect.Int {
		c.SetInt(505)
	}
}

func Set2(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)

	fmt.Println("Fields:")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ { //提取t所拥有的方法。t.NumMethod()是方法个数。
		m := t.Method(i) //m是关于方法的详细描述，包含方法的Name，Type等信息。所以m.Name是打印具体的东西。
		fmt.Println(m)
		fmt.Printf("%6s：%v\n", m.Name, m.Type) //打印Name和方法签名[func(main.User)]
		//按字母排序

		mv := v.MethodByName(m.Name)
		args := []reflect.Value{reflect.ValueOf(1)}
		mv.Call(args)
	}

}

/*z := v.FieldByName(f.Name)
		if !z.IsValid() {
			fmt.Println("不可设置")
			return
		}
		if z.Kind() == reflect.String {
			z.SetString("505")
		}
		if z.Kind() == reflect.String {
			z.SetString("已修改")
		}
	}

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("不可更改")
		return
	} else {
		v = v.Elem()//赋值这个是因为是指针所以赋值。
	}
	fmt.Println(v)

	}

	//直接输入名字--------
	f := v.FieldByName("Kaiguan")
	if !f.IsValid() {
		fmt.Println("不可设置")
		return
	}
	if f.Kind() == reflect.Int {
		f.SetInt(1001)
	}
}
*/
