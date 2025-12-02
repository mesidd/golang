package main

import "fmt"

func add(a int, b int){
	fmt.Println("Sum:",a+b)
}

func product(a int, b int, c int){
	fmt.Println("Product:",a * b * c)
}

func subtract(a, b int) int {
	return a-b
}

func main(){
	fmt.Println("Learning Function")
	add(4,5)
	product(4,9,19)
	ans := subtract(9,8)
	fmt.Println("Difference of two numbers:",ans)
}