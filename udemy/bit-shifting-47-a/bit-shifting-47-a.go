package main

import (
	"fmt"
)

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)
	//now use the bit shift operator to shift the bit count of x (4) to 1, assign to y, and print
	y := x << 1
	fmt.Printf("%d\t\t%b", y, y)
}
