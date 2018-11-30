package main

import (
	"fmt"
)

func main() {
	var a int = 1
	var p *int = &a //p指向int的指针,&a，取a的地址}

	fmt.Println(p)  //p所指向的内存地址
	fmt.Println(*p) //取出p的值
}
