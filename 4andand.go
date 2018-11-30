package main

import (
	"fmt"
)

func main() {
	a := 1
	if (10 / a) > 0 {
		fmt.Println("OK")
	}

	b := 1
	if b > 0 && (10/b) > 0 && b*b != 1 { //  &&表示如果两边都成立，才会继续执行，否则就此中断，不再继续，但也不会报错。
		//  问题：if里不能出现等号？？？？？？ 答：用 ==
		fmt.Println("NiubilitY") //因为b没有大于0，所以中断里，就打印不出NiubilitY。
	}

	if b > 0 || (10/a) > 0 { // ||或或的使用，有一个成立，则继续进行。
		fmt.Println("NIUBILITY!!!")
	}
}
