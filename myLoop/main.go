package main

import "fmt"

func main() {
	fmt.Println("Welcome with the loop in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("Index is %v and the day is %v\n", index, day)
	}

	for _, day := range days {
		fmt.Printf("The day is %v\n", day)
	}

	rougevalue := 2

	for rougevalue < 10 {

		if rougevalue == 2 {
			goto lco
		}

		if rougevalue == 5 {
			break
		}

		fmt.Println(rougevalue)
		rougevalue++
	}

lco:
	fmt.Println("Jumping at lco")

}
