package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the world of slices.")

	var fruitList = []string{"Apple", "Peach", "Orange"}

	fmt.Printf("Type of data: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])

	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	fmt.Println(highScores)

	highScores = append(highScores, 434, 34235, 63463)

	fmt.Println(highScores)

	sort.Ints(highScores)

	fmt.Println(highScores)

	//how to remove a value from slices

	var courses = []string{"reactjs", "ruby", "javascript", "Swift", "flutter", "go"}

	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
