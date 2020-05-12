package main

import (
	"fmt"
)

//"package-level" scope variables go here
//NOTICE THE TYPE IS AFTER THE VAR KEYWORD AND VARIABLE
//e.g. read as "the variable x is type int"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	//For sprintf, don't forget the verbs (i.e. format specifiers)
	//Look up the fmt package if you forget syntax, etc.
	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(s)

}

//the values of the variables printed in the output
//are referred to ZERO VALUES
