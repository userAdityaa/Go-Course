package main

import "fmt"

func main() {
	fmt.Println("Welcome to the world of maps")

	languages := make(map[string]string)

	languages["Js"] = "JavaScript"
	languages["Py"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println(languages)

	delete(languages, "RB")

	fmt.Println(languages)

	//loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("For Key %v, value is %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}

}
