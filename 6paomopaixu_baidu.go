package main

import (
	"fmt"
)

const (
	LENGTH = 8
)

func main() {
	var tmp int
	number := []int{95, 45, 15, 78, 84, 51, 24, 12}
	for i := 0; i < LENGTH; i++ {
		for j := LENGTH - 1; j > 0; j-- {
			if number[j] < number[j-1] {
				tmp = number[j-1]
				number[j-1] = number[j]
				number[j] = tmp
			}
		}
	}
	for i := 0; i < LENGTH; i++ {
		fmt.Printf("%d ", number[i])
	}
	fmt.Printf("\n")
}
