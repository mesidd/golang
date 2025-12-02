package main

import (
	"fmt"
	// "reflect"
	"strconv"
)

func main(){
	fmt.Println("Welcome to Data_Conversion!")

	// b := 10
	// c := "ram"

	// fmt.Println(reflect.TypeOf(b))
	// fmt.Println(reflect.TypeOf(c))
	// fmt.Printf("Type of b: %T\n",b)
	// fmt.Printf("Type of c: %T\n",c)

	// d := float64(b)

	// fmt.Println(b,c,d)

	// fmt.Println(reflect.TypeOf(d))

	// e := d + 1.21

	// fmt.Println(e)

	// f := 1232
	// str := strconv.Itoa(f)
	// str_int,_ := strconv.Atoi(str)
	
	// fmt.Println(str)
	// fmt.Println(reflect.TypeOf(f))
	// fmt.Println(reflect.TypeOf(str))
	// fmt.Println(reflect.TypeOf(str_int))

	g := "3.14"

	// str_float,_ := strconv.Atoi(g)
	str_float,_ := strconv.ParseFloat(g, 64)

	fmt.Println(str_float)

}