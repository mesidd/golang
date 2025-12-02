package main

import (
	"fmt"
	"os"
)


func main(){
	fmt.Println("Welcome to the File Coding!!")

	// file, _ := os.Create("file/file.txt")
	// defer file.Close()

	// content := "Hello World!"

	// // io.WriteString(file, content+"\n")

	// bytes, _ := io.WriteString(file,content+"\n")

	// fmt.Println("Size:",bytes)

	file, err := os.Open("file/file.txt")
	if err != nil{
		panic(err)
	}

	defer file.Close()

	// data, err := io.ReadAll(file)
	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Printf(string(data))

	// buffer := make([]byte, 1024)

	// for{
	// 	n, err := file.Read(buffer)
	// 	if err == io.EOF{
	// 		break
	// 	}
	// 	if err != nil{
	// 		fmt.Println("Error while reading the file:",err)
	// 		return
	// 	}

	// 	fmt.Println(string(buffer[:n]))
	// }

	content, err := os.ReadFile("file/file.txt")

	if err != nil{
		fmt.Println("Errorr",err)
		return
	}

	fmt.Println(string(content))

}