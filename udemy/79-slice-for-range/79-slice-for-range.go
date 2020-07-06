package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(len(x))
	fmt.Println(x)
	//you can access the elements by their index position
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])

	//i for index and v for value var names is commonplace
	//x is the slice you want to loop over
	//Now, range over slice x
	for i, v := range x {
		fmt.Println(i, v)
	}
}
