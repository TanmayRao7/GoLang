package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Rating for our pizza (0-5) :")

	//comma ok || error ok

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating ,", input)
	fmt.Printf("Type of rating ,%T \n", input)

}
