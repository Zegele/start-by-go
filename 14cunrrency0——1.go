package main

import (
	"fmt"
)

func main() {
	c := make(chan bool) //创建一个channel（chan bool)chan代表channel，bool代表声明类型。
	go func() {          //匿名函数
		fmt.Println("gogogo")
		c <- true //存一个布尔型true或false

	}()
	<-c //在main函数中，我们把值取出来。
	//当启动上面的goroutine后，main函数会执行到这里。
	//因为没有将内容放进去，就阻塞了，就在等，等上面输出gogogo后，将true存入channel，然后才能读出来，读出来之后，再结束运行（结束谁？）。
	//告诉main函数，我这里执行完毕了，你可以退出了。

}
