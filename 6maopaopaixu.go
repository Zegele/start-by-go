package main

import (
	"fmt"
)

func main() {
	a := [10]int{95, 60, 76, 88, 10, 23, 5, 2, 65}
	var num = len(a)
	fmt.Println(a)             //用于对比，可省
	fmt.Println(num)           //用于显示，可省
	for i := 0; i < num; i++ { //注意：=
		for j := i + 1; j < num; j++ { //注意：=
			if a[j] > a[i] { //a[i]是定的，跟所有的a[j]比较。整个历遍了后，i才i++
				temp := a[i] //为临时替换的数,注意 ：=
				a[i] = a[j]
				a[j] = temp
			}
		}

	}
	fmt.Println(a)
}
