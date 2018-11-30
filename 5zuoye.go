package main

import (
	"fmt"
)

func main() {

LABEL6:
	for i := 0; i < 5; i++ {

		for {
			fmt.Println(i)
			continue LABEL6 //contunue是带值的，去LABEL6
		}
	}
	fmt.Println("OK continue LABEL6")

LABEL7:
	for i := 0; i < 5; i++ {

		for {
			fmt.Println(i)
			goto LABEL7 //goto是不带值的去LABEL7，这就会陷入死循环。

		}
	}
	fmt.Println("OK continue LABEL7")
}
