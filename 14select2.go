//怎样让它打印10次随机数？
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {

		for i:=0; i<10;i++{
			v := range c 
			fmt.Println(v)
		}

	}()

	for {
		select {
		case c <- 0:
		case c <- 1:
		}
	}

}
