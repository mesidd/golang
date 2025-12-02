package main

import "fmt"


func modify(num *int){
	*num *= 20
}

func main(){
	fmt.Println("Welcome to Pointer Code!")

	var num int
	num = 2

	var ptr *int
	ptr = &num

	fmt.Println("Pointer:",ptr)

	var word string = "0x14000106020"

	fmt.Println("Length:",len(word))

	// var ptr2 *string
	// ptr2 = &word

	// fmt.Println("Address of word:",ptr2)

	var name string = "Ram"
	
	var ptr1 *string = &name

	number := 10

	ptr2 := &number

	fmt.Println(*ptr1,ptr1)
	fmt.Println(*ptr2,ptr2)

	var pointer *int

	fmt.Println(pointer)

	modify(ptr2)

	c := 30
	modify(&c)

	fmt.Println(*ptr2)
	fmt.Println(c)

}