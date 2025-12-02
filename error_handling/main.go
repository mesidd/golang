package main

import "fmt"

func div(a, b float64) (float64, error) {
	if b==0 {
		return 0, fmt.Errorf("denominator must not be 0")
	}
	return a/b, nil
}

func main(){
	ans, error := div(10,5)
	fmt.Println("ans:",ans, error)
}