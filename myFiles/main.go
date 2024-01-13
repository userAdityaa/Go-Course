package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to the files section")
	content := "This needs to go somewhere there."

	file, err := os.Create("./files.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("The length in the file is: ", length)
	defer file.Close()

	readFile("./files.txt")

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNilError(err)

	fmt.Println("The data in the file was: ", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
