package main

import (
	"fmt"
	"sync"
)

// 在版本1的基础上，使用循环
// 循环中 如何处理? 为什么

func main() {

	var wg sync.WaitGroup


	for _, salutation := range []string{"hello", "greetings", "good day"}{
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()





}
