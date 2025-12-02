package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){

	println("Hey, What's your name?")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Username is",name)

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("Hello, Mr.",name)
} 