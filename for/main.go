package main

import "fmt"


func main(){
	fmt.Println("Welcome to For Loop - ")

	c := 5

	for i:=0; i < c; i++{
		fmt.Printf("%d ",i)
	}
	fmt.Printf("\n")

	counter := 0

	for {
		fmt.Println("Infinite Loop")
		counter++

		if counter == 8{
			fmt.Println("Break")
			break
		}
	}

	numbers := []int {1,2,3,4,5}

	for index, value := range numbers{
		fmt.Printf("Index: %d, value: %d\n", index, value)
	}

	data := "Hello, World!"

	for i, char := range data{
		fmt.Printf("%d -> %c\n", i, char)
	}
}