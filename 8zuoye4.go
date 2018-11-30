package main

import (
	"fmt"
)

func main() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	m2 := make(map[string]int)

	for k, v := range m1 {
		m2[v] = k //v不会影响m1，但没说不会影响m2啊！！！！
		fmt.Println(m2[v])
		fmt.Println(m2) //这个结果告诉我们已经起作用了。
	}

}
