package main

import (
	"fmt"
)

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	y = 911
	fmt.Printf("%#x\n", y)
	//this line, each y gets it's own formatter
	//and we use \t for tabs instead of \n for new line (just another option)
	fmt.Printf("%#x\t%b\t%x", y, y, y)

	// string print func and assign to variable
	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Println(s)
	//for the value in the default format
	fmt.Printf("%v", y)

}
