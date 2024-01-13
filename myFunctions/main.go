package main

import "fmt"

func main() {
	fmt.Println("Welcome to the main function ")
	greeter()
	result := adder(3, 5)
	fmt.Println(result)

	fmt.Println(proAdder(1, 2, 3, 4, 5, 1, 2, 3, 5))

	number, value := message(33)

	fmt.Println(value)
	fmt.Println(number)
}

func adder(one int, two int) int {
	return one + two
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

func message(values ...int) (int, string) {
	return 34, "Aditya Chaudhary"
}

func greeter() {
	fmt.Println("Hello World")
}
