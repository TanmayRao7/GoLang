package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	const url = "https://enterprise.dev.hotelkeyapp.com/v4/version"

	fmt.Println("Web Requests in Go Lang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Response is of type %T\n", response)

		defer response.Body.Close()

		databytes, err := ioutil.ReadAll(response.Body)

		if err != nil {
			panic(err)
		}

		content := string(databytes)

		fmt.Println("The contents of file are \n", content)

	}

}
