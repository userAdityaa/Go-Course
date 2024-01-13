package main

import "fmt"

func main() {
	// defer fmt.Println("Yeh baad mein aega")
	// fmt.Println("Hello Defer")

	defer fmt.Println("Pehla")
	defer fmt.Println("Dusra")
	defer fmt.Println("Teesra")
	fmt.Println("Yeh toh hona hi tha")
}
