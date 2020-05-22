package main

import (
	"fmt"
)

const (
	_ = iota
	// kb = 1024 //a kb (in binary) is 1 shifted over 10 times
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t \t%b\n", gb, gb)

	//bit shifting is a little esoteric but worth knowing. Despite being perhaps
	//overly concise, the instructor believes the code above is still idiomatic.
}
