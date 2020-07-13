package main

import (
	"fmt"
)

type person struct {
	last  string
	first string
	age   int
}

func main() {
	p1 := person{
		last:  "Bond",
		first: "James",
		age:   32,
	}
	fmt.Println(p1)
}
