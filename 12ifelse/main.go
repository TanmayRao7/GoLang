package main

import "fmt"

func main() {

	fmt.Println("If Else in Go lang")

	loginCount := 9

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount >= 10 {
		result = "Not a regular user"
	} else {
		result = "Not Defined User"
	}

	fmt.Println(result)

}
