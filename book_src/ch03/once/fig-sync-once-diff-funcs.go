package main

import (
	"fmt"
	"sync"
)

// 问题， 下面的输出是什么

func main() {
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)

	fmt.Printf("Count: %d\n", count)
}
