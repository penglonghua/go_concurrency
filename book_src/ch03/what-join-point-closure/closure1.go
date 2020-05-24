package main

import (
	"fmt"
	"sync"
)

// 问题1 ， 闭包可以从创建它们的作用域中获取变量。
// 那么闭包是在这些变量的副本上运行，还是原始值的引用上运行?
// 这里输出的是什么?

func main() {

	var wg sync.WaitGroup

	salutation := "hello"

	wg.Add(1)

	go func() {
		defer wg.Done()
		salutation = "welcome" // 这个地方只有一个值，这是肯定的.

	}()

	wg.Wait()
	fmt.Println(salutation)   // 输出的是 welcome



}
