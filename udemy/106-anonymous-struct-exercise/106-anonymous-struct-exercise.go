package main

import (
	"fmt"
)

//note I didn't create a named struct type here, I just stick
//a unnamed (i.e. anonymous) type under main

func main() {
	//create anonymous struct using a composite literal
	b1 := struct { //have to assign it to a var identifier and use the identifier for this to work
		fork       string
		rearTravel int
	}{
		fork:       "Fox",
		rearTravel: 150,
	}

	fmt.Println(b1)

}
