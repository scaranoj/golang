package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	fmt.Println(x[0:5])  //you don't have to specify the 0 as the low here, can just omit
	fmt.Println(x[5:10]) //you don't have to specify the 10 as the high here, can just omit
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])

}
