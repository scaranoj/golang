package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("The modulus of %d is %d when divided by 4\n", i, i%4)
	}
}
