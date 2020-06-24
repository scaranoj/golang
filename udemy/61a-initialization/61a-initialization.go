package main

import (
	"fmt"
)

func main() {
	if x := 42; x == 42 {
		fmt.Println("001")
	}
	//note that the scope of x would prevent the below statement from running
	//fmt.Println(x)
}
