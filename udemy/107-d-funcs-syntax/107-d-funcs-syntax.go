package main

import (
	"fmt"
)

func main() {
	foo()
	bar("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	//Below is a call to the mouse func. "Ian" and "Fleming" are the 2 inputs for the mouse func that I'm calling,
	//and I'm assigning the outputs to x and y (i.e. you're calling the mouse func
	//and passing it 2 string inputs, Ian and Fleming, which returns back to you a
	//string reply and a boolean value of true)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)

}

// This is the syntax our funcs will have:
// func (r receiver) identifier(paramaters) (return(s)) { ... }

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(st string) string {
	return fmt.Sprint("Hello from woo,", st)
}

//here's a func with multiple returns
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprintln(fn, " ", ln, `, says "Hello"`)
	b := true //we just picked this for the sake of showing multiple returns from a func
	return a, b
}
