package main

import (
	"fmt"
)

func main() {
	foo()
}

// This is the syntax our funcs will have:
// func (r receiver) identifier(paramaters) (return(s)) { ... }

func foo() {
	fmt.Println("hello from foo")
}

//the "hello from foo" is considered a "parameter"
//the value in the parens of the func call on line 8 is considered an "argument"
