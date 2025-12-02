package main

import "fmt"

func main(){
	x := 10

	if x > 5{
		fmt.Println("X > 5")
	} else{
		fmt.Println("X is not > 5")
	}

	y := 45
	if y > 10{
		fmt.Println("y > 10")
	} else {
		fmt.Println("y < = 10")
	}

	fmt.Println("Enter the value to be checked")

	var z int
	
	fmt.Scan(&z)

	if z % 2 == 0{
		fmt.Println("Even")
	}else {
		fmt.Println("Odd")
	}

}