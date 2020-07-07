package main

import (
	"fmt"
)

func main() {
	// a slice of string as a composite literal
	jb := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	//create a slice of a slice of string type, and assign it to xp
	xp := [][]string{jb, mp}
	fmt.Println(xp)
	//notice that it's basically just a composite literal slice and that...
	//the entire type is made up of the 2 sets of brackets plus the type name.
	//Then the curly braces are for the values, which in this case are
	//variables that refer to other slices....which begins to create rows of a psudo-spreadsheet

}
