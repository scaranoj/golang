package main

import (
	"fmt"
)

type salami int

//remember the order when declaring a variable with it's type
//variable first, then it's type (i.e. "the variable x of type salami")

var x salami

func main() {
	//guess what the value will be here, and what is it called?
	//Answer: 0, "zero value"
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	//the variable x is at package-level scope so this func sees it
	x = 42
	fmt.Println(x)
}
