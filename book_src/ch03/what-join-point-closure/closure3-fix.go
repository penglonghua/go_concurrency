package main

import (
	"fmt"
	"sync"
)

// 如何在循环中正确处理这个,
// 在循环中 传入副本

func main() {

	var wg sync.WaitGroup

	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()

}
