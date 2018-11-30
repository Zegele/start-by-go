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
	Kaiguan int
}

func (u ShiZhong) DaoShu(TingZhi int) {
	i := u.Count
	switch {
	case i > TingZhi:
		for i := u.Count; i > TingZhi; i-- {
			fmt.Println(i)
		}
	case i == TingZhi:
		fmt.Println(i)
	default:
		for i := u.Count; i < TingZhi; i++ {
			fmt.Println(i)
		}
	}
}

func main() {
	Plan1 := ShiZhong{Name: "miqi", Count: 10, KaiGuan: KaiGuan{Kaiguan: 0}}

	Plan1.DaoShu(20)

	v := reflect.ValueOf(Plan1)
	mv := v.MethodByName("DaoShu")
	args := []reflect.Value{reflect.ValueOf(20)}
	mv.Call(args)

	Plan2 := Plan1.KaiGuan
	fmt.Println(Plan2)
	Set1(&Plan2)
	fmt.Println(Plan2)
	fmt.Println(Plan1)

	Plan3 := ShiZhong{Name: "miqi", Count: 100, KaiGuan: KaiGuan{Kaiguan: 1}}

	Set2(Plan3)
	fmt.Println(Plan3)

	Set3(&Plan3)
	fmt.Println(Plan3)

}

//--------------改值
func Set1(q interface{}) {
	v := reflect.ValueOf(q)
	fmt.Println(v)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("不可更改")
		return
	} else {
		v = v.Elem()
	}
	fmt.Println("h", v)

	f := v.FieldByName("Kaiguan")
	fmt.Println(f)
	if !f.IsValid() {
		fmt.Println("不可设置")
		return
	}
	if f.Kind() == reflect.Int {
		f.SetInt(1001)
	}

	u := v.FieldByName("Name")
	fmt.Println(u)
	if !u.IsValid() {
		fmt.Println("不可设置")
		return
	}
	if u.Kind() == reflect.String {
		u.SetString("已修改")
		fmt.Println(u)
	}

	c := v.FieldByName("Count")
	fmt.Println(c)
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

	fmt.Println("Type:", t.Name())
	fmt.Println("这里Fields:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println("nnn", f)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)

		z := v.FieldByName(f.Name) //怎样通过历遍名字修改值呢？？？
		fmt.Println("zzzz", z)

	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m)
		fmt.Printf("%6s：%v\n", m.Name, m.Type)
		mv := v.MethodByName(m.Name)
		fmt.Println(mv)
		args := []reflect.Value{reflect.ValueOf(90)}
		mv.Call(args)
	}
}

func Set3(q interface{}) {
	t := reflect.TypeOf(q)
	fmt.Println("Type:", t.Name())
	v := reflect.ValueOf(q)
	fmt.Println(t)

	fmt.Println(v)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("不可更改")
		return
	} else {
		v = v.Elem()
	}

	fmt.Println(v.NumField())

	for i := 0; i < v.NumField(); i++ {
		fmt.Println("hhhh", v.Field(i))

	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s：%v\n", m.Name, m.Type)
		mv := v.MethodByName(m.Name)
		args := []reflect.Value{reflect.ValueOf(90)}
		mv.Call(args)
	}
}
