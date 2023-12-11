package main

import "fmt"

func main() {

	fmt.Println("Structs in golang")

	robo := User{"Robot", "example@gmail.com", true, 16}
	fmt.Println(robo)
	fmt.Printf("Robots Details are: %+v\n ", robo)
	fmt.Printf("Robots Name are: %v\n ", robo.Name)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
