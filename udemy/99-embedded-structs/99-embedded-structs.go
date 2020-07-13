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

type secretAgent struct {
	person //notice how we didn't give this field [name?] a value, apparently it's not necessary
	ltk    bool
}

//notice how we can mix diff types (e.g. strings and ints)
//which you can't do with arrays, slices, and maps
func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true, //the compiler formatter did this syntax weirdness - FYI
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(p2.first, p2.last, p2.age)
}
