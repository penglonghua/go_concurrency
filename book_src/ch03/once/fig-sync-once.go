package main

import (
	"fmt"
	"sync"
)

// 问题， 下面的输出是什么?

func main() {
	var count int

	increment := func() {
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment) // 貌似 once只有一个 Do方法
		}()
	}

	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}
