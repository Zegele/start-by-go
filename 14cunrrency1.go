package main

import (
	"fmt"
)

func main() {
	c := make(chan bool) //创建一个channel（chan bool)chan代表channel，bool代表声明类型。
	go func() {          //匿名函数
		fmt.Println("gogogo")
		c <- true //存一个布尔型true或false//执行完的时候是一个存的操作。

		fmt.Println("222gogogo")
		close(c)
		fmt.Println("111111gogogo")
	}()
	for v := range c { //这里在不断的存取，迭代c(channel)。
		fmt.Println(v) //在进行跌代的时候，它一直在等channel有个值存进去，然后再取出来，然后再打印出来。
	}
	//goroutine是goroutine，channel是channel。
}
