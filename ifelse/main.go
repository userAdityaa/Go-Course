package main

import "fmt"

func main() {
	fmt.Println("If else hain bhai")

	loginCount := 23
	var result string

	if loginCount < 18 {
		result = "smaller user"
	} else if loginCount == 23 {
		result = "Miracle Miracle"
	} else {
		result = "normal user"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Yeh toh sahi hogaya na toh ")
	} else {
		fmt.Println("Yeh toh galat hogaya yaarr")
	}

}
