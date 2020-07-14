package main

import (
	"fmt"
)

type person struct {
	first       string
	last        string
	favIceCream []string
}

func main() {
	p1 := person{
		first:       "James",
		last:        "Bond",
		favIceCream: []string{"Rocky Road", "Whisky", "Lemon"},
	}
	p2 := person{
		first:       "Miss",
		last:        "Moneypenney",
		favIceCream: []string{"Strawberry", "Vanilla", "Hazelnut"},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favIceCream {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favIceCream {
		fmt.Println(i, v)
	}

	fmt.Printf("\n*** Here's a nested for ***\n\n")

	//here's a loop within a loop for giggles
	for i, v := range p1.favIceCream {
		fmt.Println(i, v)
		for i, v := range p2.favIceCream {
			fmt.Println(i, v)
		}

	}

}
