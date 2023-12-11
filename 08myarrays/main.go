package main

import "fmt"

func main() {

	fmt.Println("Arrays in Go lang ")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Peach"

	fmt.Println("Fruit List = ", fruitList)
	fmt.Println("Fruit List = ", len(fruitList))

	var vegList = [3]string{"1", "2", "3"}
	fmt.Println(vegList)
	fmt.Println(len(vegList))

}
