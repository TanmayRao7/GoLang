package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON !")
	//EncodeJson()
	DecodeJSON()
}

func EncodeJson() {

	lcoCourses := []course{
		{"ReactJs Bootcamp", 299, "lco.dev", "123@", []string{"web-dev", "js"}},
		{"Mern Bootcamp", 399, "lco.dev", "123@", []string{"web-dev", "front-end"}},
		{"Angular Bootcamp", 199, "lco.dev", "123@", []string{"google", "js"}},
	}

	//package this data as JSOn data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJSON() {

	jsonDataFromWeb := []byte(`
	{
		"coursename": "Mern Bootcamp",   
		"price": 399,
		"website": "lco.dev",
		"tags": ["web-dev","front-end"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid !")
	}

	//add data to key:value pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", lcoCourse)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v\n", k, v)
	}

}
