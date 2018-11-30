package main

import (
	"fmt"
)

type A struct {
	B
}

type B struct {
	C
	Name string
}

type C struct {
	Name string
}

func main() {
	b := B{C: C{"C"}, Name: "B"}
	a := A{b}
	fmt.Println(b)
	fmt.Println(b.Name, b.C.Name)
	fmt.Println(a)
	fmt.Println(a.Name, a.B.Name) //由于B和C的Name同名，在A中只找最近一级的同名Name
}
