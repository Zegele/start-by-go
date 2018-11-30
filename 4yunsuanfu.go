package main

import (
	"fmt"
)

func main() {
	fmt.Println(2 ^ 2)   //^二元运算符
	fmt.Println(^2)      //^一元运算符,右边有数字
	fmt.Println(!false)  //!取反
	fmt.Println(1 << 10) //1左移10
	fmt.Println(1 << 10 << 10)
	fmt.Println(1 << 20) //1左移10,左移10，右移10
	fmt.Println(2 << 10)
	/*
		    6: 0110
		   11: 1011
		   ----------
		   &   0010  = 2
		   如果两个是1则是1，否则是0

		   |   1111  =15
		   如果由一个是1，则是1，否则为0

		   ^   1101  =13
			有一个为1，则为1，两个为1，则为0
			*问题：两个为0，会怎样？答：依然为0


		   &^  0100  =4
		   第二位数（即1011），如果有一位是1，则会将第一位数（即0110）的相应位数强制改为0，否则不变。
	*/
	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 &^ 11)

	/*
		6: 0110
		8: 1000
		----------
		&  0000=0
		|  1110=14
		^  1110
		&^ 0110=6
	*/
	fmt.Println(6 & 8)
	fmt.Println(6 | 8)
	fmt.Println(6 ^ 8) //^的二元运算
	fmt.Println(6 &^ 8)

}
