package main

import "fmt"

func main() {
	//nesting loops
	//the first for (i) is the outer loop, and the second for (j) is the inner
	for i := 0; i <= 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("ther outer loop has run %d\t The inner loop %d\n", i, j)
		}
	}
}
