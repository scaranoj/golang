package main

import (
	"fmt"
)

func main() {
	favSport := "football"
	switch favSport {
	case "basketball":
		fmt.Println("this shouldn't print")
	case "baseball":
		fmt.Println("this shouldn't print")
	case "tennis":
		fmt.Println("this shouldn't print")
	case "football":
		fmt.Println("correct! football is my fav sport!")
	default:
		fmt.Println("that is not my favorite sport, sorry!")

	}
}
