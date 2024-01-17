package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to the real backend")
	// performGetRequest()
	performPostRequest()
}

func performPostRequest() {
	const myUrl = "http://localhost:3000"

	requestBody := strings.NewReader(`
		{
			"coursename" : "Let's go with golang", 
			"price": 0, 
			"platform": "learnCodeOnline.in" 
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func performGetRequest() {
	const myUrl = "http://localhost:3000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic((err))
	}

	response.Body.Close()

	fmt.Println("The status code is", response.StatusCode)
	fmt.Println("The content length is", response.ContentLength)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(content))

	byteCount, _ := responseString.Write(content)

	fmt.Println("The byte count is:", byteCount)

	fmt.Println(responseString.String())
}

func performPostFormRequest() {
	const myUrl = "http://localhost:3000"

	data := url.Values{}
	data.Add("firstname", "Aditya")
	data.Add("lastname", "Chaudhary")

	reponse, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer reponse.Body.Close()

	content, _ := ioutil.ReadAll(reponse.Body)

	fmt.Println(content)

}
