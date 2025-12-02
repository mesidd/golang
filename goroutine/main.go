package main

import (
	"fmt"
	"time"
)

func sayHello(){
	fmt.Println("Hello World")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Ended First Function!")
}

func sayHi(){
	fmt.Println("Hi Sid!")
	time.Sleep(1000* time.Millisecond)
}

func main(){
	// fmt.Println("Welcome to Goroutine")
	go sayHello()
	go sayHi()

	time.Sleep(3000 * time.Millisecond)
}

