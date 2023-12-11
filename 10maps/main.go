package main

import "fmt"

func main() {

	fmt.Println("Maps in goLang")

	language := make(map[string]string)

	language["Java"] = "Java11"
	language["Java1"] = "Java8"
	language["Java2"] = "Java13"

	fmt.Println("Maps = ", language)
	fmt.Println(language["Java"])

	delete(language, "Java")
	fmt.Println("Maps = ", language)

}
