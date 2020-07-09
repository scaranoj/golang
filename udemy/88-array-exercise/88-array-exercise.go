package main

import (
	"fmt"
)

func main() {
	//create a composite literal array
	a := [5]int{
		1, 2, 3, 4, 5,
	}

	for i := range a {
		fmt.Println(i)
	}
	fmt.Printf("%T", a)

}
