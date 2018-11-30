package main

import "fmt"

var (
	name  = "liuzexian"
	name2 = "LIUZEXIAN"
)

func main() {
	var a, b, c, _, d int = 1, 2, 3, 4, 5 //这种方式就是并行方式,_为空白符号，用于忽略
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	var e, f = "1", 2 //省略了 string、int
	fmt.Println(e)
	fmt.Println(f)

	g, h := false, 2.1 //省略了var，但别忘了是:=。因为：相当于var
	fmt.Println(g)
	fmt.Println(h)

	var i float32 = 1.12345678 //小数后第八位有时被舍去，有时被进一位,其实已经超出里float32的范围
	j := int(i)                //演示里将浮点型转化为整型，这里的：相当于var
	fmt.Println(i)
	fmt.Println(j)
}
