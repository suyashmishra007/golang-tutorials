package main

import (
	"fmt"
	"sync"
)

func worker(i int,wg *sync.WaitGroup){
	defer wg.Done(); // Signal that the goroutine end.
	fmt.Printf("worker %d started\n",i);
	// some task is happening
	fmt.Println("Task end",i)
}

func main() {
	var wg sync.WaitGroup;
	for i:= 1 ; i <= 3 ; i++ {
		wg.Add(1); // Increment the waitGroup counter.
		go worker(i,&wg);
	}
	wg.Wait();
	fmt.Println("worker Task completed");
}
