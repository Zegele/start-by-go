package main

import (
	"fmt"
	"time"
)

func main() {
	go Go()
	time.Sleep(2 * time.Second)
	//go和mail都在运行，但mail函数sleep后，Go开始执行，然后退出程序。
	//如果没有Sleep，main直接就停止了，是不会管Goroutine是不是继续。
}

func Go() {
	fmt.Println("Go GO GO!!!")
}
