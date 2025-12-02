package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup){
	fmt.Printf("Worker task started %d\n", i)
	// task
	fmt.Printf("Worker task ended %d\n", i)
	defer wg.Done()
}


func main(){
	fmt.Println("Welcome to SyncWaitGroup")

	// start 3 worker goroutines
	var wg sync.WaitGroup

	for i:=1; i <= 3; i++{
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("Worker task finished!")
}