package main

import (
	"fmt"
)

func main() {
	for {
		years_alive := 1989
		for {
			if years_alive <= 2020 {
				fmt.Println(years_alive)
				years_alive++
			}
		}
	}
}
