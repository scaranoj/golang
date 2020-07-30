package main

import (
	"fmt"
)

func main() {
	x := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The total is", x)
}

//Remember, the syntax for a func is:
// func (r receiver) identifier(parameter) (return(s)) { code }

//here, we create a variadic parm with the `...` indicating that we want an unlimited number of int's

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", v)
	}
	fmt.Println("The total is", sum)
	return sum
}
