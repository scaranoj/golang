package main

import (
	"fmt"
)

//just get rid of the type declaration header that was originally placed here
//and use it's contents (values) within the struct instance, which creates an
//"anonymous" struct

func main() {
	p1 := struct {
		last  string //notice how we didn't create the struct type above? This is totally legal in go
		first string
		age   int
	}{
		last:  "Bond",
		first: "James",
		age:   32,
	}
	fmt.Println(p1)
}
