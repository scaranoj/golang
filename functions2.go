//when two or more function parms share the same type, you
//can omit the type from all but the last

package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
