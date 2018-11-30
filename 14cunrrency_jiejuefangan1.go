//稳定输出10个的解决方案一
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10) //存入10个
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	fmt.Println("aaa") //为什么打印的aaa、bbb也没有规律可循？尝试理解吧，重视结果就好。。。

	for i := 0; i < 10; i++ {
		<-c
	}
	fmt.Println("bbb")
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 4; i++ {

		a += i
	}
	fmt.Println(index, a)

	c <- true
}
