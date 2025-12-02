package main

import "fmt"


func main(){
	// fmt.Println("Welcome to Defer Tutorials!!")
	defer fmt.Println("1 - End of the Program")
	fmt.Println("2 - Middle")
	defer fmt.Println("3 - Start of the Program")
	fmt.Println("4 - Cross of the Program")
	defer fmt.Println("5 - Intermediate")
}