package main

import (
	"fmt"
)

//create TYPED and UNTYPED constants. Print the values of the constants

const (
	a = 41
	b = 42
)

//typed - don't forget to put the type AFTER the variable...ALWAYS
const c int = 43
const d int = 44

func main() {
	fmt.Println(a, b, c, d)
}
