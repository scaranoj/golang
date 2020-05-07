package main

import "fmt"

//DECLARE the variable y and ASSIGN the value
//of 43
//declare and assign = initialization
var y = 43

//DECLARE that there is a variable with IDENTIFIER "z"
//and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
//ASSIGNS the ZERO VALUE of TYPE int to "z"

var z int

func main() {
	// short declaration operator
	//DECLARE a variable and ASSIGN a VALUE for a certtain type
	x := 42
	fmt.Println(x)

	fmt.Println(y)

	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println("again:", y)

}
