package main

import (
	"fmt"
)

var y string
var z int

func main() {
	//Declare a variable to be a certain type
	//Assign a VALUE of that type to the variable
	fmt.Println("Hello Colorado!")

	//printing the value of y
	fmt.Println("printing the value of y", y, "<-- any value there after y?")
	fmt.Printf("%T\n", y)

	y = "Bond, James Bond"

	fmt.Println("printing the value of y", y, "<-- any value there after y?")
	fmt.Printf("%T", y)

	fmt.Println("printing the value of z", z, "<-- any value there after z?")
	fmt.Printf("%T", z)
}
