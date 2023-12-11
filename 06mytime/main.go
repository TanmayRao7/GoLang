package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcom to Time study of Go lang")

	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.October, 10, 23, 23, 0, 0, time.Local)

	fmt.Println(createdDate)

	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

}

//https://communityday.awsugkochi.in/
