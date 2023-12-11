package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Welcome to files in Go Lang")
	content := "This needs to go in a file - www.google.com"

	file, err := os.Create("./contents.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length is :", length)

	defer file.Close()

	readFile("./content.txt")

}

func readFile(filename string) {

	databyte, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside the file is \n", string(databyte))

}
