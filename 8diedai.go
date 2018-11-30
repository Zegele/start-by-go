package main

import (
	"fmt"
)

func main() {
	/*
		迭代：for配合range  用于slice或map

		for i, v:=range slice{//迭代slice；i表示slice的索引是int型；v表示slice索引应对的值，是对slice值的复制，不会影响slice。
		slice[i] = ____ //如果要改变slice的值，可以用这种方法。
		}

		for k,v:=range map{//迭代map：k表示key,v表示value，是对map值的拷贝的操作，不对map的值造成影响。

		}
	*/

	sm := make([]map[int]string, 5, 10)
	fmt.Println(sm)
	for _, v := range sm { //不需要索引，就用下划线代替表示。
		v = make(map[int]string, 2) //这里的2是什么意思？cap容量的意思。
		v[0] = "NIUBILITY"
		v[1] = "OK"

		v[3] = "OK？" //超过cap,会重新分配内存地址。
		fmt.Println(v)
		fmt.Printf("%p\n", v) //这个地址就没看懂了，不管超或不超，地址都在变。
		fmt.Printf("%p\n", sm)
	}
	fmt.Println(sm) //因为v是对拷贝进行操作，所以对原sm是没有影响的。

	//---------------------------将改变的值给到原slice

	for i := range sm {
		sm[i] = make(map[int]string) //初始化
		sm[i][1] = "NIUBILITY!"
		fmt.Println(sm[i]) //sm[i]是个map,这个懂的吧。。。,map是无序的。
	}
	fmt.Println(sm)

	//---------------------------------

}
