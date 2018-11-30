package main

import (
	"fmt"
)

const (
	ZERO = iota
	BITE
	KB
	MB
	TB
)

func main() { //别把mian忘了
	if BITE == 1 {
		fmt.Println("bite")
		fmt.Println(BITE)
	}

	if (BITE << (BITE * 10)) == 1024 {
		fmt.Println("KB")
		fmt.Println(BITE << (BITE * 10))
	}

	if (BITE << (BITE * 10 * KB)) == (1024 * 1024) {
		fmt.Println("MB")
		fmt.Println(BITE << (BITE * 10 * KB))
	}

	if (BITE << (BITE * 10 * MB)) == (1024 * 1024 * 1024) {
		fmt.Println("TB")
		fmt.Println(BITE << (BITE * 10 * MB))
	}

}
