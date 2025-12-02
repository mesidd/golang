package main

import (
	"fmt"
	"myproject/myutil"
)

func main(){
	fmt.Println("Hello world")
	myutil.PrintMessage("This is go language.")

	var age int = 10
	var name string = "Sidd"
	fmt.Println(name, "Age: ", age)

	var roll string = "22PE10030"
	fmt.Println(name, age, roll);

	person := "Ram"
	person = "Shyam"
	fmt.Println(person)


}