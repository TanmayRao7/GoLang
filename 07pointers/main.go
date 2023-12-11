package main

import "fmt"

func main() {

	fmt.Println("Pointers")

	// var ptr *int
	// fmt.Println("Value of ptr = ", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of ptr = ", ptr)
	fmt.Println("Value of ptr = ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New Value is: ", myNumber)

}
