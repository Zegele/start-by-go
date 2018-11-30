package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i) //4-1
		//对参数i的拷贝
	}

	//---------------------------------------

	for e := 0; e < 3; e++ {
		defer func() { //在main函数为turn的时候，我们开始执行defer语句
			fmt.Println(e) //3-1
			//引用的局部变量的i,引用i的地址，i已经是结束循环后的3了
		}() //注意()别忽略！！！因为defer是在调用某个函数。
	}
	//---------------------------------------------
	//探索
	//把这两个合起来，先参数后函数
	for i := 0; i < 3; i++ {
		defer fmt.Println(i) //2-2
		defer func() {
			fmt.Println(i) //2-1
		}()
	}

	//----------------------------------------------
	//把这两个合起来，先函数后参数
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i) //1-1
		}()
		defer fmt.Println(i) //1-1
	}
	fmt.Println("sheixian?")
}
