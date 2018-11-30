package main

import (
	"fmt"
)

var c chan string

func Pingpong() {
	i := 0
	for {
		fmt.Println(<-c)
		c <- fmt.Sprintf("From Pingpong:Hi,#%d", i) //Sprintf是什么意思？？？
		//Sprintf将输出的东西，都是字符串类型的。例如这个i就转化成字符串类型输出了。
		i++
	}
}

func main() {
	c = make(chan string)
	go Pingpong()
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("From Main: Hello,#%d", i) //这个值写入后，便被上面读取了，所以下面就变成了等待写入，再读取。
		fmt.Println(<-c)                            //因为被上面的读了所以执行到这里又在等待，写入，写入后，才能被读取。
	}
}
