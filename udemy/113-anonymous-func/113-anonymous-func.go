package main

import (
	"fmt"
)

func main() {
	foo()

	func() {
		fmt.Println("Anonymous func ran")
	}() // note the empty parens just outside the anon func body? That's it!
	// Those parens is what makes it execute
	// we don't need to pass in an argument since it takes no argument

	// here is a func where we can pass in arguments
	func(x int) {
		fmt.Println("The meaning of life is", x)
	}(42) // here we're passing in an argument
}

func foo() {
	fmt.Println("Hello Colorado!")
}
