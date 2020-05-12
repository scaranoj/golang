package main

import (
	"fmt"
)

//"package-level" scope variables go here
//NOTICE THE TYPE IS AFTER THE VAR KEYWORD AND VARIABLE
//e.g. read as "the variable x is type int"

var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

//the values of the variables printed in the output
//are referred to ZERO VALUES
