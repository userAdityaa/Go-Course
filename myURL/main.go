package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://loc.dev:3000/learn?coursename=reactjs&paymentid=dfwrkls"

func main() {
	fmt.Println("Welcome to the handling URL")

	fmt.Println(myURL)

	//parsing
	result, _ := url.Parse(myURL)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("The type of query params is %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, value := range qparams {
		fmt.Println("Param is: ", value)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=aditya",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
