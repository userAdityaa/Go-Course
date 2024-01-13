package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Struct in golang")

	aditya := User{"Aditya", "aditya.chaudhary1558@gmail.com", true, 18}

	fmt.Println(aditya)

	fmt.Printf("Aditya all details are here: %+v\n", aditya)

	fmt.Printf("Name of the user: %v\n", aditya.Name)
}
