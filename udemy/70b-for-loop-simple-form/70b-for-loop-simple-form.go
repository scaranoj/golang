package main

import (
	"fmt"
)

func main() {
	for {
		yearsAlive := 1989
		for {
			if yearsAlive <= 2020 {
				fmt.Println(yearsAlive)
				yearsAlive++
			}
		}
	}
}
