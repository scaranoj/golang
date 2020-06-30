package main

import (
	"fmt"
)

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%U\n", i)
		}
	}
}

//first attempt below. I was trying to start
//with the unicode char set while the instructor
//just used the associated decimal number with the requested char set

/*
func main() {
	for x := 'A'; x <= 'Z'; x++ {
		fmt.Printf("%d\n", x)

		for r := 'A'; r <= 'Z'; r++ {
			fmt.Printf("/t/t%U\n/t/t%U\n/t/t%U\n", r, r, r)

		}
	}

}
*/
