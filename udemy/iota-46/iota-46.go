package main

import (
	"fmt"
)

const (
	a = iota + 1
	b = iota
	c = iota
)

//or you could do this

const (
	d = iota + 1
	e = iota
	f = iota
)

//or you could also do this

const (
	g = iota
	h
	i
)

//keep in mind that the series will reset to 0 every time you declare const

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

}
