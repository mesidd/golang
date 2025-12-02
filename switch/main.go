package main

import "fmt"

func main(){
	
	fmt.Println("Welcome to Switch Cases!")

	fmt.Printf("Enter the value: ")

	var z int

	fmt.Scan(&z)

	switch z {
	case 1,10:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Out of index!")
	}

}