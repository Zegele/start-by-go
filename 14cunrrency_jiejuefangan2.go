//稳定输出10个的解决方案二
//任务组是不是都完成，待完成任务减为0时，就结束了。
package main

import (
	"fmt"
	"runtime"
	"sync" //同步
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{} //这种方法就不用创建channel，创建WaitGroup
	wg.Add(10)             //增加任务数
	for i := 0; i < 10; i++ {
		go Go(&wg, i) //指针的拷贝,传进去指针的地址。
		//如果是(wg, i)会有这个错误提示：  cannot use wg (type sync.WaitGroup) as type *sync.WaitGroup in argument to Go
	}
	fmt.Println("aaa") //为什么打印的aaa、bbb也没有规律可循？尝试理解吧，重视结果就好。。。

	wg.Wait()
	fmt.Println("bbb") //这个aaa，bbb还是不可预测的啊。。。 看结果就好。。。试着理解吧。。。
}

func Go(wg *sync.WaitGroup, index int) {
	//如果是(wg sync.WaitGroup, index int)会有这个错误提示： cannot use &wg (type *sync.WaitGroup) as type sync.WaitGroup in argument to Go
	a := 0
	for i := 0; i < 4; i++ {

		a += i
	} //斐波那契数列
	fmt.Println(index, a)

	wg.Done() //完成任务就标记1次，任务数消去一个。
}
