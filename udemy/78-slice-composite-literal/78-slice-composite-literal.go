package main

import (
	"fmt"
)

//below is an example of a composite literal slice
//the syntax is `x := type{values}`
func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	fmt.Println(x[4])
}
