package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	//note that this again is a composite literal and that
	//the entire expression after the assignment operator is the type
	//(i.e. map[string]int{}
	fmt.Println(m)
	fmt.Println(m["James"])

	//what happens when you check for a key that's not in the map?
	fmt.Println("Barnabas") //returns 0

	//check to see if the key "Barnabas" is in the map using the
	// "comma ok" golang idiom.

	v, ok := m["James"]
	fmt.Println(v)
	fmt.Println(ok)

	f, cool := m["Barnabas"]
	fmt.Println(f)
	fmt.Println(cool)

	//you can make this ok check into an if statment
	if v, ok := m["Barnabas"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)
	}

	//create another map on your own using map/composite literal
	m = map[string]int{
		"Age":    32,
		"Weight": 180,
	}
	//check for a value in the o map above
	w, good := m["Age"]
	fmt.Println(w, good)

	if w, good := m["Weight"]; good {
		fmt.Println(w, good)
	}

}

//use an if statement to determine if value is present while
//also checking for a successful lookup result
