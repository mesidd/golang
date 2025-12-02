package main

import (
	"fmt"
	"net/url"
	"reflect"
)

func main(){
	fmt.Println("Welcome to the URL Coding")

	URL := "https://jsonplaceholder.typicode.com/todos/1?key1=value1"
	// URL := "ram@123"
	fmt.Println(reflect.TypeOf(URL))

	parsed_url, err := url.Parse(URL)
	if err != nil{
		fmt.Println("Error",err)
		return
	}
	// fmt.Println("Type of parsed_url:",reflect.TypeOf(parsed_url))

	fmt.Println("Scheme:",parsed_url.Scheme)
	fmt.Println("Host:",parsed_url.Host)
	fmt.Println("Path:",parsed_url.Path)
	fmt.Println("RawQuery:",parsed_url.RawQuery)

	parsed_url.Path = "/todos/2"
	parsed_url.RawQuery = "username:mesidd"

	newURL := parsed_url.String()

	fmt.Println(newURL)

}

