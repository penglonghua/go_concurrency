package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {


	var wg sync.WaitGroup

	wg.Add(1) // +1
	go func() {
		defer wg.Done() // -1
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()

	wg.Add(1)  // +1
	go func() {
		defer wg.Done() // -1
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()

	wg.Wait() // 等待
	fmt.Println("All goroutines completes.")

}
