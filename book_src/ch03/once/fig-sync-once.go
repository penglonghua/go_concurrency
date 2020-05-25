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
	increments.Add(100) //  100个 go

	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()

			once.Do(increment) // 貌似 once只有一个 Do方法, 这个地方提示的只有方法
			// 说明, Do这里，传递进来的是 函数, 即使在不同的 go 上面，也只会调用一次 once.Do(fun)
		}()
	}

	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}
