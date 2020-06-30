package main

import (
	"fmt"
)

func main() {
	x := 1
	if x > 5 {
		fmt.Printf("%d is greater than 5", x)
	} else if x < 5 {
		fmt.Printf("%v is less than 5", x)
	} else {
		fmt.Println("x is 5!")
	}
}
