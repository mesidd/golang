package main

import (
	"encoding/json"
	"fmt"
)

type Person struct{
	Name string `json:"name"`
	Age int `json:"age"`
	IsAdult bool `json:"is_adult"`
}

func main(){
	fmt.Println("Welcome to JSON Golang")

	person := Person{Name: "Ram",Age:21,IsAdult: true}

	fmt.Println("Person:",person.Age)

	// convert person into json Encoding (Marshaling)

	jsonData, err := json.Marshal(person)
	if err != nil{
		fmt.Println("Error:",err)
		return
	}

	fmt.Println("JSON_DATA:", string(jsonData))

	var personData Person

	err = json.Unmarshal(jsonData, &personData)
	if err != nil{
		fmt.Println("Error:",err)
		return
	}
	fmt.Println("Data:",personData)
}