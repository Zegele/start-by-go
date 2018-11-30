package main

import (
	"fmt"
)

func main() {
	var fs = [4]func(){} //这是一个以四个函数为元素的切片。

	for i := 0; i < 4; i++ { //i循环四次

		defer fmt.Println("defer i =", i) //打印参数值//第三打印，这里的i是作为参数的。
		//--------------------

		defer func() {
			fmt.Println("defer_closure i=", i) //第二打印，引入参数地址，是最终的值
		}()
		//--------------------

		fs[i] = func() {
			fmt.Println("closure i =", i) //对i没有定义，没有引用参数，就夺取了i的地址
		} //这是在定义，还没调用，没有打印（参照niminghanshu.go)
		//--------------------
	}

	for _, f := range fs {
		f() //对这个函数进行调用，引用地址
		/*f()是func(){
			fmt.Println("closure i =", i)
		}
		*/

	}
	//触发main的时候才会触发defer!!!!!!!!!!!!!!!
}

/*猜测值为：
closure i = 4
closure i = 4
closure i = 4
closure i = 4
defer_closure i=4
defer i =3
defer_closure i=4
defer i =2
defer_closure i=4
defer i =1
defer_closure i=4
defer i =0
*/
