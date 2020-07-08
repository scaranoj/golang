package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)
	fmt.Println(m["James"])

	//add an element
	m["Jason"] = 33

	//range over to see all the elements in the map
	for k, v := range m {
		fmt.Println(k, v)
	}

}
