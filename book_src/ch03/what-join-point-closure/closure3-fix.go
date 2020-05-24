package main

import (
	"fmt"
	"sync"
)

// 如何在循环中正确处理这个,
// 在循环中 传入副本

// 输出的顺序是不固定的，这个地方还体现出了 并发的作用
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
