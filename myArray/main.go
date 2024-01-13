package main

import "fmt"

func main() {
	fmt.Println("Welcome to the array list.")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"

	fmt.Println("The fruit list is:", fruitList)
	fmt.Println("The length of the fruit array is:", len(fruitList))
	fmt.Println("")
}
