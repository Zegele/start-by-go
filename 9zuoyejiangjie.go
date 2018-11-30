package main

import (
	"fmt"
)

func main() {
	var fs = [4]func(){} //这是一个以四个函数为元素的切片。

	for i := 0; i < 4; i++ { //i循环四次

		defer fmt.Println("defer i =", i) //打印参数值//第三打印，这里的i是作为参数的。值拷贝。
		//--------------------

		defer func() {
			fmt.Println("defer_closure i=", i) //第二打印，引入参数地址，是最终的值
		}()
		//--------------------

		fs[i] = func() {
			fmt.Println("closure i =", i) //这个是匿名函数。函数体内并没有定义i，所以i是从外层函数中获取的，所以它夺取的是i的引用地址。执行完for循环后i其实就等于4了，所以再调用输出这个i就是4了。这个地址对应的值就已经是4。
		}

	}

	for _, f := range fs {
		//在上个for循环中将匿名函数存到了以fs[i]类型（function函数类型）的slice中
		//在这个for循环之后，再用一个for循环调用这个函数,来输出值。
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
