//怎样让它打印10次随机数？实现了
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	for i := 0; i < 10; i++ {

		select {
		case c <- 0:
		case c <- 1:
		}
	}
}
