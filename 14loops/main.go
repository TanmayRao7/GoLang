package main

import "fmt"

func main() {

	fmt.Println("Welcome to loops in Go Lang")

	days := []string{"Sun", "Mon", "Tue", "Wed"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, value := range days {
		fmt.Printf("Index is %v and value is %v \n", index, value)
	}

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 5 {
			//break
			rogueValue++
			continue
		}
		if rogueValue == 2 {
			goto lco
		}

		fmt.Println("Value is: ", rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("Jumping at a random line")

}
