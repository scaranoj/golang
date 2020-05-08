package main

import (
	"fmt"
)

var a int

//create your own type
type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	//the next 4 lins wouldn't work because
	//you can't assign type hotdog to an int
	//go is statically typed and won't allow it
	//once you declare a type, you can't change it
	//a = b
	//fmt.Printf("%T\n", a)
	//fmt.Println(b)
	//fmt.Printf("%T\n", b)

}
