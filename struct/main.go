package main

import "fmt"

type Person struct {
	FirstName string
	LastName string
	Age int
}

func main(){
	// fmt.Println("Welcome to Struct Class!")

	p1 := Person{"Ram", "Roy", 21}

	// fmt.Println("Details of P1:")
	// fmt.Printf("Name: %s %s \nAge: %d \n",p1.FirstName, p1.LastName, p1.Age)

	var p2 Person

	p2.FirstName = "Ganesh"
	p2.LastName = "Lakshmi"
	p2.Age = 20

	// fmt.Printf("%s %s is %d years older. \n", p2.FirstName, p2.LastName, p2.Age)

	var p3 = new(Person)
	
	p3.FirstName = "Mahesh"
	p3.LastName = "Rana"
	p3.Age = 25

	p4 := Person{
		FirstName: "Sidd",
		LastName: "Roy",
		Age: 24,
	}
	
	// fmt.Println(p1,p2,p3)
	fmt.Printf("p1: %v\np2: %v\np3: %v\n", p1, p2, *p3)
	fmt.Printf("p4: %s %s %d\n", p4.FirstName, p4.LastName, p4.Age)
	
}