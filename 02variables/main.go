package main

import "fmt"

const LoginToken string = "qwrt" // public

func main() {

	//String
	var username string = "TanmayRao7"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	//String
	var isTrue bool = true
	fmt.Println(isTrue)
	fmt.Printf("Variable is of type: %T \n", isTrue)

	//unint
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	//unfloat
	var smallfloat float32 = 255.23243245557674
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type: %T \n", smallfloat)

	//deafault values and aliases
	var anontherVariable int
	fmt.Println(anontherVariable)
	fmt.Printf("Variable is of type: %T \n", anontherVariable)

	// implicit type
	var website = "www.google.com"
	fmt.Println(website)

	// no var style
	numberOfUser := 30000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
