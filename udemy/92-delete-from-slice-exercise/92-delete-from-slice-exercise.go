package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	//need to delete 45, 46, and 47 by slicing out and appending to a new slice

	l := append(x[:3])
	m := append(x[6:])

	y := append(l, m...)
	fmt.Println(y)

	//OR, a faster way w/o having to introduce variables and less lines (that
	//the instructor used would be:
	// x = append(x[:3], x[6]...)
	// fmt.Println(x)

}
