package main

import "fmt"

func main() {

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	//add an element
	m["Bob"] = 34

	//now print the map with the new element
	fmt.Println(m)

	//delete an element
	delete(m, "Bob")

	//now print the map one more time with the new element removed
	fmt.Println(m)

	//write an if lookup on the map and delete the key "James"

	if v, ok := m["James"]; ok {
		fmt.Println(v)
		delete(m, "James")
	}
	fmt.Println(m)
}
