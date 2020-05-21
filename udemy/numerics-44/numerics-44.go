package main

import (
	"fmt"
)

func main() {
	// assign the string H to a variable s and print it out
	s := "H"
	fmt.Println(s)

	// convert the string s to a byte slice and assign, it to the variable bs, and print it out
	bs := []byte(s)
	fmt.Println(bs)

	//select the 0 index of the bs slice and declare/assign to the variable n
	n := bs[0]
	fmt.Println(n)
	//print the type of n
	fmt.Printf("%T\n", n)
	//print the binary value of n
	fmt.Printf("%b\n", n)
	//print the hexidecimal value of n
	fmt.Printf("%#x\n", n)
}
