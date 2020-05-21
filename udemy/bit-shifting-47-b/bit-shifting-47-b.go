package main

import (
	"fmt"
)

func main() {
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t \t%b\n", gb, gb)

	// notice how the binary digit shifted over (<<) another 10 bits (i.e. places) each subsequent time
	//becuase of it shifting 10 each time, this is a great use case for an iota (see bit-shifting-47-c.go)
}
