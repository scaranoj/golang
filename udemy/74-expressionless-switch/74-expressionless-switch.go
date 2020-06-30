package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("this shoudn't print because it's false")
	case (2 == 4):
		fmt.Println("should print false")
	case (4 == 4):
		fmt.Println("4 is equal to 4, and therefore this should print!")
	}
}
