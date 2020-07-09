package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	//fmt.Println(x, y)
	xy := [][]string{x, y}
	fmt.Println(xy)

	for i, xs := range xy {
		fmt.Println("record: ", i)
		for j, xxs := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v\n", j, xxs)
		}

	}

}
