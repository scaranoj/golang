package main

import (
	"fmt"
)
//notice I didn't bother creating a new struct type here?
//it's because I'm only using a struct in one little area
//and want to keep this code lightweight and clean and want to avoid 
//creating extraneous types where I don't need them
//therefore, I'm just creating an anonymous struct under main

func main() {
	p1 := struct {
		last  string
		first string
		age   int
	}{
		last:  "Bond",
		first: "James",
		age:   32,
	}
	fmt.Println(p1)
}
