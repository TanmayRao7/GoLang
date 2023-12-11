package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	fmt.Println("Welcome to Web Server")

	//PerformGetRequest()
	//performPostJSONRequest()
	performPostFormRequest()

}

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte Count is: ", byteCount)
	fmt.Println(responseString.String())

}

func performPostJSONRequest() {

	const myurl = "http://localhost:3000/post"

	requestBody := strings.NewReader(`
		{
			"coursename":"Go Lang",
			"price":100,
			"platoform":"lco"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}

func performPostFormRequest() {

	const myurl = "http://localhost:3000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "Tanmay")
	data.Add("lastname", "Rao")
	data.Add("Email", "raotanmay97@gmail.com")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
