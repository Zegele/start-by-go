package main

import (
	"fmt"
)

type TZ int

func (a *TZ) Increase(num int) { //a这里的a就是一个符号。可以换成其他，比如tz,不影响调用或结果。
	*a += TZ(num) //使用TZ类型的num，对num进行强制化转化。 a+=b----a=a+b
	//因为a是TZ类型的，num是int类型的，两个是不同类型的，所以，不能直接运算。所以把num转化成TZ类型的，才能进行计算。
}

func main() {
	var a TZ
	a = 0
	a.Increase(11111111) //这里输入参数，这里输入的值直接默认是int型
	fmt.Println(a)
}

/*

package main

import (
	"fmt"
)

type TZ int

func (tz *TZ) Increase(num int) {  //这里的tz和上面的a都可以用，不影响调用。
	*tz += TZ(num) //使用TZ类型的num，对num进行强制化转化。 a+=b----a=a+b
	//因为a是TZ类型的，num是int类型的，两个是不同类型的，所以，不能直接运算。所以把num转化成TZ类型的，才能进行计算。
}

func main() {
	var a TZ
	a = 0
	a.Increase(11111111) //这里输入参数，这里输入的值直接默认是int型
	fmt.Println(a)
}

*/
