package main

import (
	"fmt"
)

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	// can't do this next line becauase can't assign one type to another type
	// a = b
	//You can however, do this which CONVERTS the value of b (type hotdog) to an int,
	// which will  allow you to store a value of int into a previously declared var of TYPE hotdog
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
