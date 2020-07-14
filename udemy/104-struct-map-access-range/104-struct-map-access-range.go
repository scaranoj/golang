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

	//store the values of type person in a map with the key of last name

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for i, v := range m {
		fmt.Println(i)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for k, val := range v.favIceCream {
			fmt.Println(k, val)
		}
	}

}
