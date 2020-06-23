package main

import (
	"fmt"
)

//print the remainders up to 100
func main() {
	x := 1
	for {
		//the increment has to be her at top
		//instead of bottom becuase the compiler
		//will never get to it because it will
		//keep breaking out and jumping to top on 1
		x++
		if x > 100 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
}
