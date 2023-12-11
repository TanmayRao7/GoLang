package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Introduction to Slices")

	var fruitList = []string{"Tomato", "Apple", "Onion"}

	fmt.Printf("The type of = %T \n ", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println("fruitList =", fruitList)

	fruitList = append(fruitList[1:])

	fmt.Println("fruitList =", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 1
	highScores[1] = 4
	highScores[2] = 3
	highScores[3] = 2

	highScores = append(highScores, 8, 9)

	sort.Ints(highScores)

	fmt.Println(highScores)

	// How to remove a value from slices

	var courses = []string{"react", "next", "python", "jquery"}
	fmt.Println("Courses = ", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
