package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(x)

	for v, i := range x {
		fmt.Println(v, i)
	}
	fmt.Printf("%T", x)
}
