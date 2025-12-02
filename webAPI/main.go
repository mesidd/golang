package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	// "reflect"
)


func main(){
	fmt.Println("Welcome to API coding!")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil{
		fmt.Println("Error",err)
		return
	}

	defer res.Body.Close()
	// fmt.Println("Type:", reflect.TypeOf(res))

	data, err := ioutil.ReadAll(res.Body)

	if err != nil{
		fmt.Println("Error:",err)
		return
	}
	fmt.Println("Response:",string(data))
}