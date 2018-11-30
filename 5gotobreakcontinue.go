package main

import (
	"fmt"
)

func main() {

	//break 标签------------------------------------------
LABEL1: //定义标签1，注意 ‘：’冒号
	for {
		for i := 0; i < 5; i++ { //理解没？i=0初始值先进入if,不符合后走到i++（1），再进入if不符合if,继续i<4,i++，并且i不可能==4
			if i > 3 {
				break LABEL1 //break LABRL1直接结束标签，如果不加标签，只能结束该级的循环（可能会造成无限循环）。
				//break 标签   理解为跳出某个标签
			}
			fmt.Println(i)
		}
	}
	fmt.Println("OK break LABEL1")

	//goto 标签---------------------------------------------

	for {
		for i := 0; i < 5; i++ { //理解没？i=0初始值先进入if,不符合后走到i++（1），再进入if不符合if,继续i<5,i++
			if i > 3 {
				goto LABEL2 //如果LABEL2在上面，有可能会出现死循环。所以一般放下面。
				//goto 标签     理解为进入某个标签。
			}
			fmt.Println(i)
		}
	}
LABEL2:
	fmt.Println("OK goto LABEL2")

	//continue 无标签  表示不进行剩下的循环体，直接进行下一次循环。
	//continue 标签--------------------------------------
LABEL3:
	for i := 0; i < 5; i++ {
		for {
			fmt.Println(i)
			continue LABEL3 //表示继续进行LABEL3
			//注意这里的LABEL3便不能是无线循环形式。
		}
	}
	fmt.Println("OK continue LABEL3")

LABEL4:
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		for {
			continue LABEL4 //表示继续进行LABEL3
			//注意这里的LABEL3便不能是无线循环形式。
		}
	}
	fmt.Println("OK continue LABEL4")

LABEL5:
	for i := 0; i < 5; i++ {
		for {
			continue LABEL5 //表示继续进行LABEL3
			//注意这里的LABEL3便不能是无线循环形式。
			//	fmt.Println(i) 这句放在这里也无法执行，因为continue又已经返回去执行LABEL5了
		}
		fmt.Println(i) //这句放在这里也无法执行，因为continue又已经返回去执行LABEL5了
	}
	fmt.Println("OK continue LABEL5")
} //这个括号怎么有时候就自动删除了？这句放在这里也无法执行，因为continue又已经返回去执行LABEL5了
