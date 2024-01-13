package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello Switch Please")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of the dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You can move 1")
	case 2:
		fmt.Println("You can move 2")
	case 3:
		fmt.Println("You can move 3")
	case 4:
		fmt.Println("You can move 4")
		fallthrough
	case 5:
		fmt.Println("You can move 5")
	case 6:
		fmt.Println("You can move 6")
	}
}
