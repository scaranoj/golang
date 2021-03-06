package main

import (
	"fmt"
)

//create a new struct type called person and create 3 fields called first, last, and age
type person struct {
	first string
	last  string
	age   int
}

//notice how we can mix diff types (e.g. strings and ints)
//which you can't do with arrays, slices, and maps
func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
