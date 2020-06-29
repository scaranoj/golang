package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("this should not print because it's false")
	case (2 == 4):
		fmt.Println("this should not print because it's false")
	case (3 == 3):
		fmt.Println("print")
		//the fallthrough keyword will allow the next expression to be evaluated
		fallthrough
	case (4 == 4):
		fmt.Println("also true, does it print?")
		fallthrough
	case (5 == 5):
		fmt.Println("also true, does it print?")
		fallthrough
	default:
		fmt.Println("this is default")
	}
}
