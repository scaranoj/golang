package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	//assign some values to some of the positions and print them out
	x[0] = 42
	x[9] = 999

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	//try to assign a value to the 10 index position
	//you'll receive an "index out of range" error
	//x[10] = 4005
	//fmt.Println(x)

	//but you can change the slice length with the built-in append func
	x = append(x, 4005)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 4006)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 4007)
	//notice from the output how the go runtime decided to double the cap to accomodate
	//the new value at index position 13
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
