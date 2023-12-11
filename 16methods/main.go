package main

import "fmt"

func main() {

	fmt.Println("Structs in golang")

	robo := User{"Robot", "example@gmail.com", true, 16}
	fmt.Println(robo)
	fmt.Printf("Robots Details are: %+v\n ", robo)
	fmt.Printf("Robots Name are: %v\n ", robo.Name)
	robo.GetStatus()
	robo.NewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {

	fmt.Println(u.Status)
	fmt.Println(u.Name)
	fmt.Println(u.Age)
	fmt.Println(u.Email)

}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of user: ", u.Email)
}
