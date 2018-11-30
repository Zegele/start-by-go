package main

import (
	"fmt"
)

func main() {
	a := 1
	//switch 第一种用法 switch后跟表达式
	switch a { //当a a算是条件表达式，如果switch没有条件表达式，要在case后放条件表达式
	case 0: //      =0时，执行下面
		fmt.Println("a=0")
	case 1: //      =1时，执行下面
		fmt.Println("a=1")
	default: //      都不是时，执行下面
		fmt.Println("None")
	}

	//switch 第二种用法 switch后无表达式，case后跟表达式
	switch {
	case a >= 0: //表示：如果 a >= 0，则执行下面。
		fmt.Println("a>=0") //如果没有"fallthrough"，符合一个case,则自动结束。
	case a >= 1: //没有fallthrough，后面的就算符合也不会继续执行了。
		fmt.Println("a>=1")
	default:
		fmt.Println("None")
	}

	//switch 第二种用法 switch后无表达式，case后跟表达式，并且有fallthrough
	switch {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough //fallthrough表示：如果符合，也继续下一个case。
	case a >= 1:
		fmt.Println("a>=1")
	default:
		fmt.Println("None")
	}

	//switch 第三种用法 switch后 跟初始化语句
	switch b := 2; { //将 b:=2 该初始表达式放在里switch后，但只在switch语句块内部有效。
	//注意：  ；分号
	case b >= 1:
		fmt.Println("b>=1")
		fallthrough //fallthrough表示：如果符合，也继续下一个case。
	case b >= 2:
		fmt.Println("b>=2")
	default:
		fmt.Println("None")
	}
}
