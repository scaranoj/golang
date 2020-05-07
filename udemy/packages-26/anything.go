package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Println("Sup homeboy?", false, 24)
	fmt.Println(n)
	fmt.Println(err)
}
