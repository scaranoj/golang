package main

import (
	"fmt"
)

func main() {
	foo()
	bar("James")
}

// This is the syntax our funcs will have:
// func (r receiver) identifier(paramaters) (return(s)) { ... }

func foo() {
	fmt.Println("hello from foo")
}

// everything in Go is PASS BY VALUE
//we're passing the value "James" into string
func bar(s string) {
	fmt.Println("Hello,", s)
}

//the "hello from foo" is considered a "parameter"
//the value in the parens of the func call on line 8 is considered an "argument"
