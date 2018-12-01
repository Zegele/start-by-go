package anzhuangbao //包名和文件夹名最好保持一致

import (
	"fmt"
)

func Hello() {
	fmt.Println("Hello from anzhuangbao")
	hello() //调用hello函数，开头小写只能被内部调用
}

func hello() { //这个是小写开头的，不能被外部调用。
	fmt.Println("hello from anzhuangbao???")
}
