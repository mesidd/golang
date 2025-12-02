package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func performGetRequest(){
		res, err := http.Get("https://jsonplaceholder.typicode.com/todos/100")
	if err != nil{
		fmt.Println("Error:",err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK{
		fmt.Println("Error in fetching response:", res.Status)
		return
	}

	// data,error := ioutil.ReadAll(res.Body)
	// if error != nil{
	// 	fmt.Println("Error:",error)
	// 	return
	// }
	// fmt.Println("Data:",string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil{
		fmt.Println("Error:",err)
	}
	fmt.Println("Todo:",todo)

	fmt.Println("comp:",todo.Completed)
}
type Todo struct{
	UserID int `json: user_id`
	ID int `json: id`
	Title string `json: title`
	Completed bool `json: completed`
}

func performPostRequest(){
	todo := Todo{
		UserID: 101,
		Title: "Sid",
		Completed: true,
	}

	myURL := "https://jsonplaceholder.typicode.com/todos/"

	// convert todo to json
	jsonData, err := json.Marshal(todo)
	if err != nil{
		fmt.Println("Error", err)
	}

	// convert json to string
	jsonString := string(jsonData)

	// convert string to reader
	jsonReader := strings.NewReader(jsonString)
	
	// send POST Request
	res, err := http.Post(myURL, "application/json",jsonReader)
	if err != nil{
		fmt.Println("Error:",err)
		return
	}
	
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil{
		fmt.Println("Error:",err)
	}

	fmt.Println("Status:",res.Status)
	fmt.Println("Response:", string(data))

}

func performUpdateRequest(){
	todo := Todo{
		UserID: 11,
		Title: "Sid Is Here",
		Completed: false,
	}

	// convert todo to json
	jsonData, err := json.Marshal(todo)
	if err != nil{
		fmt.Println("Error:",err)
	}

	myURL := "https://jsonplaceholder.typicode.com/todos"

	// json string
	jsonString := string(jsonData)

	// json reader
	jsonReader := strings.NewReader(jsonString)
	
	// create PUT Request
	req, err := http.NewRequest(http.MethodPut, myURL, jsonReader)
	if err != nil{
		fmt.Println("Error:",err)
	}
	req.Header.Set("Content-type","application/json")

	// client sent the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil{
		fmt.Println("Error:",err)
		return
	}
	
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response: ",string(data))

	fmt.Println("Status:", res.Status)
}

func performDeleteRequest(){

	const myURL = "https://jsonplaceholder.typicode.com/todos"

	// create Delete Request
	req, err := http.NewRequest(http.MethodDelete, myURL, nil)
	if err != nil{
		fmt.Println("Error:",err)
		return
	}

	// send the request
	client := http.Client{}
	res, err := client.Do(req)

	if err != nil{
		fmt.Println("Error:",err)
	}

	defer res.Body.Close()

	// data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Status:",res.Status)
	
}

func main(){
	fmt.Println("Welcome to CRUD Operations")
	// performGetRequest()
	// performPostRequest()
	// performUpdateRequest()
	performDeleteRequest()

}
