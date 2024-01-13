package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Welcome to methods")

}

func (u User) GetStatus() {
	fmt.Println("The user is active: ", u.Status)
}
