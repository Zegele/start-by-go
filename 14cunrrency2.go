package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("GOGOGO!!!")
		c <- false
	}()
	<-c //取channel的值，但没有东西可以取，等上面的gotoutine中给channel传入false后，才有值可以被取。

	d := make(chan bool, 1) //1代表缓存数
	go func() {
		fmt.Println("222GOGOGO!!!")
		d <- true
	}()
	<-d //有缓存时，d channel已经输入了值，所以后面，这里读出d的值，然后关闭channel

	go func() {
		fmt.Println("222GOGOGO!!!")
		<-d
	}()
	d <- true //有缓存时（d channel的缓存是1），写入值后，至于你读不读无关。所以上面的读出值时，已经关闭了。
}

//如果是无缓存，取的动作要在存的前面。无缓存是同步阻塞的。
//如果是无缓存,必须有读出才能停止，一直没有读出一直会阻塞，执行中（等待）。
//如果是有缓存，是异步的
