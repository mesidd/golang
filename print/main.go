package main

import "fmt"

func main(){
	name := "Sid"
	age := 24
	role := "Software Engineer"

	// fmt.Println("Name",name,"Age",age,"Role",role)

	// fmt.Printf("Name - %s\n",name)
	// fmt.Printf("Age - %d\n",age)
	// fmt.Printf("Hey, I used printf.\n")
	// fmt.Printf("Type of name is %T\n", name)

	fmt.Printf("Name: %s\n",name,"Age: %d\n",age,"Role: %s\n",role)

	// fmt.Printf("Name: %s\nAge: %d\nRole: %s\n", name, age, role)

}