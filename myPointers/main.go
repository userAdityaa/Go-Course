package main

import "fmt"

func main() {
	fmt.Println("Hello Pointers")

	// var ptr *int

	// fmt.Println("Value of the pointer is:", ptr)

	myNumber := 32

	var ptr = &myNumber

	fmt.Println("The value of pointer is:", ptr)
	fmt.Println("The value of actual pointer is:", *ptr)
}
