package main

import "fmt"

func main() {
	//loop from 33 to 122 to print the numbers
	//and then show them as turn into the letters using format printing
	num := 33
	for num <= 122 {
		fmt.Printf("%d %c\n", num, num)
		num++
	}

}
