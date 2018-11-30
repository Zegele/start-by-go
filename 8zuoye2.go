package main

import (
	"fmt"
)

func main() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	m2 := make(map[string]int)
	i := 0
	for k, v := range m1 {
		if i == k {
			temp := k
			m2[i] = temp

			fmt.Println(m2[v])
		}
		i++
	}
	fmt.Println(m2)
}
