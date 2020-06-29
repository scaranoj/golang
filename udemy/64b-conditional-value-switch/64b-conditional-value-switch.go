package main

import (
	"fmt"
)

func main() {
	switch "Bond" {
	case "Moneypenny":
		fmt.Println("this is money")
	case "Q":
		fmt.Println("this is Q")
	case "Bond":
		fmt.Println("James Bond")
	default:
		fmt.Println("this is default")
	}
}
