package main

import (
	"fmt"
)

func main() {
	x := 44
	if x == 40 {
		fmt.Println("Our value was 40")
	} else if x == 41 {
		fmt.Println("Our value was 41")
	} else if x == 42 {
		fmt.Println("Our value was 42")
	} else if x == 43 {
		fmt.Println("Our value was 43")
	} else {
		fmt.Println("Our value was not 40, 41, 42, or 43")
	}

}
