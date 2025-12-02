package main

import "fmt"

type Person struct {
	FirstName string 
	LastName string
	Age int
}

type Contact struct{
	Email string
	Phone string
}

type Address struct{
	House int
	Area string
	State string
}

type Employee struct{
	person_details Person
	contact_details Contact
	address_details Address
}


func main(){
	
	fmt.Println("Welcom to Nested Struct")

	var p1 Employee

	p1.address_details = Address{11, "Vizag", "India"}
	p1.person_details = Person{"Ram", "Dinkar", 25}
	p1.contact_details = Contact{"ram@gmail.com", "987"}

	fmt.Printf("Person: %v\nContact: %v\nAddress: %v\n",p1.person_details, p1.contact_details, p1.address_details)
	fmt.Println("Age: ",p1.person_details.Age)

}