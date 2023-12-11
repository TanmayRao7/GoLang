package main

import "fmt"

func main() {

	fmt.Println("Welcome to functions in GoLang")
	greeter()
	greeter2()
	result := adder(3, 5)
	fmt.Println(result)

	proRes := proAdder(2, 5, 8, 7, 1, 2, 3, 45)
	fmt.Println(proRes)

}

func greeter() {
	fmt.Println("Namaste from GoLang")
}
func greeter2() {
	fmt.Println("Another Method")
}
func adder(val1 int, val2 int) int {

	return val1 + val2

}
func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total

}
