package main

import (
	"fmt"
	// "reflect"
	"time"
)

func main(){
	fmt.Println("Welcome to Time in Golang!!!\n")

	currentTime := time.Now()

	// fmt.Println("Time:", currentTime)
	// fmt.Println("Time Type:",reflect.TypeOf(currentTime))

	// // formatted_data := currentTime.Format("dd-mm-yyyy")
	// formatted_data := currentTime.Format("02-01-2006")
	// fmt.Println(formatted_data)

	// formatted_data = currentTime.Format("2006-01-02")
	// fmt.Println(formatted_data)
	
	// formatted_data = currentTime.Format("2006-01-02, Monday, 03:04 AM" )
	// fmt.Println(formatted_data)

	// layout_str := "2006 -01-02"
	// date_str := "2025-12-02"
	// formatted_time, _ := time.Parse(layout_str, date_str)

	// fmt.Println(formatted_time)
	// fmt.Println("Type of formatted_time:",reflect.TypeOf(formatted_time))
	// fmt.Println("Type of formatted_time:",reflect.TypeOf(date_str))
	// fmt.Println("Type of formatted_time:",reflect.TypeOf(layout_str))

	new_date := currentTime.Add(24 * time.Hour)
	fmt.Println(new_date)

	formatted_new_data := new_date.Format("2 January 2006 Monday")
	fmt.Println(formatted_new_data)


}