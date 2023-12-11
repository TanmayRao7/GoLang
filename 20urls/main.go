package main

import (
	"fmt"
	"net/url"
)

func main() {

	const myurl string = "https://api.github.com/users/hadley/orgs"

	fmt.Println("Handling Urls in Go Lang")

	fmt.Println("Getting data from ->", myurl)

	//parsing
	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query param is %T\n", qparams)

	fmt.Println(qparams)

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)

}
