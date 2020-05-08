package main

import (
	"fmt"
	"sync"
)

// 其中一种 join point

func main() {


	var wg sync.WaitGroup

	sayHello := func() {
		wg.Done()
		fmt.Println("hello")
	}

	wg.Add(1)
	go sayHello()
	wg.Wait()


}




