package main

import (
	"fmt"
)

const x = "Hola Colorado!"

//you can also group them together with parens, similar to the above import statement
const (
	a = 42
	b = 68
)

//notice how the above constants were untyped (i.e no type was assigned)
//you can explicitly assign types too if you want

const y int = 77

func main() {
	fmt.Println(x)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
}
