package main

import (
	"fmt"
)

func main(){
	fmt.Println("Welcome to Slice")

	// numbers := []int{1,2,3,4,5}

	// numbers = append(numbers,8,9,11,23,12)

	// fmt.Println("Numbers:",numbers)
	// fmt.Printf("Type of Numbers: %T\n",numbers)
	// fmt.Println("Length of Array:", len(numbers))

	// number := []int{1,2,3,4}
	// number := []int{}

	number := make([]int, 3, 5)

	number = append(number, 4,5,6,7)
	number = append(number, 5,6,7)
	number = append(number, 11)

	// number = append(number, 8,9,9,1,2)

	fmt.Println("Slice:",number)
	fmt.Println("Size:",len(number))
	fmt.Println("Capacity:",cap(number))

}