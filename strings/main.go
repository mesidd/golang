package main

import (
	"fmt"
	"strings"
)

func main(){
	
	fmt.Println("Welcome to Strings!!")

	// data := "apple, orange, banana"
	// parts := strings.Split(data, ",")
	// fmt.Println(parts)

	// dot_data := "ram.shyam.palak"
	// parts_name := strings.Split(dot_data, ".")
	// fmt.Println(parts_name)

	// str := "one two one two one four five"
	// var count int = strings.Count(str,"one")

	// fmt.Println("Count of 'one':",count)


	// str = "  Hello, World!   "
	// trimmed := strings.TrimLeft(str," ")

	// fmt.Printf(trimmed)
	// fmt.Println(str)

	first := "Ram"
	last := "Gopal"

	result1 := strings.Join([]string{first, last}," ")
	result2 := strings.Join([]string{first, last}," and ")

	fmt.Println(result1)
	fmt.Println(result2)
}