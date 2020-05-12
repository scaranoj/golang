package main

import (
	"fmt"
)

//Don't put the variables here, you'll get an error that
//the vars weren't created inside the func main (which they should be)

func main() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}
