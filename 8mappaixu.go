package main

import (
	"fmt"
	"sort" //排序包
)

func main() {
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make([]int, len(m))
	fmt.Println(s)
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	sort.Ints(s) //slice是引用类型
	fmt.Println(s)
}
