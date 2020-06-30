package main

import (
	"fmt"
)

// use a for conditional statement to print the years you've been alive

func main() {
	bd := 1985
	for bd <= 2020 {
		fmt.Println(bd)
		bd++

	}
}

//below was my first attempt
/*
func main() {
	age := 0

	for age <= 29 {
		fmt.Println(age)
		age++
	}
}
*/
