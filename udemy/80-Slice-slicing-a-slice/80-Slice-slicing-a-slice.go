package main

import (
	"fmt"
)

func main() {
	//create a slice and slice it with the colon operator
	//"slicing it" just means to return the values at the
	//beginning and ending index positions you specified
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[1:3])
	fmt.Println(x[2:4])

	for i, v := range x {
		fmt.Println(i, v)
	}
	//Here's how to loop w/o using a range - just as an exercise
	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
}
