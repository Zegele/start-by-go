package main

import (
	"fmt"
)

const a int = 1
const b = "A" //定义常量方法一

const ( //方法二
	c = a
	d = a + 1
	e = a + 2
)

const E, f = 2, "4" //方法三 注意先后的 逗号

const ( //方法四
	g, h = "1", 10
)

const ( //方法五
	i = 1
	j //没有给初始值，就使用上行的值
	k
)

var L string = "123"

const (
	l = "123"
	m = len(l) //内置函数才行，如果用L就不行。
	n
)

const (
	o, p = 1, "2"
	q, r //必须对应出现
)

func main() {
	fmt.Print(a)
	fmt.Print(b)
	fmt.Print(c)
	fmt.Print(d)
	fmt.Print(e)
	fmt.Print(E)
	fmt.Print(f)
	fmt.Print(g)
	fmt.Print(h)
	fmt.Print(i)
	fmt.Print(j)
	fmt.Print(k)
	fmt.Print(l)
	fmt.Print(L)
	fmt.Print(m)
	fmt.Print(n)
	fmt.Print(o)
	fmt.Print(p)
	fmt.Print(q)
	fmt.Print(r)
}
