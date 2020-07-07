package main

import (
	"fmt"
)

func main() {
	//create a slice and slice it with the colon operator
	//"slicing it" just means to return the values at the
	//beginning and ending index positions you specified
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 114, 1014)
	fmt.Println(x)

	y := []int{234, 456, 789}
	x = append(x, y...)
	fmt.Println(x)

	//if you wanted to delete the 7 and 8 in the slice, then run the below
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

}
