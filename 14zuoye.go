package main

import (
	"fmt"
)

func main() {
	c1, c2 := make(chan string), make(chan string)
	o := make(chan bool) //信号通道，用于检验c1,c2是否执行完了。
	go func() {
		for { //使用无限循环，因为如果不使用无限循环，只会执行一次select
			select {
			case v, ok := <-c2: //ok parten???
				if !ok { //如果ok不等于true的话 //当channel结束的话，会退出select。
					o <- true //读到true，说明c1取完了,然后到最后，读出o，然后，main关闭。
					break
				}
				fmt.Println("c1", v)

			case v, ok := <-c1: //如果channel是关闭的，这个ok就是false。虽然关闭了，但ok还是在读取c2
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()

	c2 <- "hi"
	c1 <- "有人吗"
	c1 <- "能好好说话吗"
	c2 <- "不能"

	close(c1)
	close(c2)

	<-o //bool的channel，只要读出true或false，这个channel就关闭了，结束了。

}
