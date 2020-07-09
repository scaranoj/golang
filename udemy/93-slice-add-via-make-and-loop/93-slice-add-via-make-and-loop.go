package main

import "fmt"

func main() {
	states := make([]string, 50, 500)
	fmt.Println(len(states))
	fmt.Println(cap(states))

	//put a value into each of the 50 positins in the length of the slice
	//we are putting values into the 50 positions that are the length of the slice
	for i := 0; i < 0; i++ {
		states[i] = fmt.Sprintf("Position %d", i)
	}

	fmt.Println(states)
	fmt.Println(len(states))
	fmt.Println(cap(states))

	//append the values to the slice grows the length of the slice
	//the underlying array "capacity" of 500 is the same

	states = append(states, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)

	fmt.Println(states)
	fmt.Println(len(states))
	fmt.Println(cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])
	}

}
