package main

import (
	"fmt"
	"sync"
)

//  这个地方作者的经验
// 只调用一次 add ,来追踪一组 goroutine
// 也可以在 for循环调用 add,接着是 go .

func main() {


	// 函数参数的 值指针指针
	hello := func(wg *sync.WaitGroup, id int) {
		defer  wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}

	const numGreeters = 5
	var wg sync.WaitGroup
	wg.Add(numGreeters)  // 调用一次, 一下增加了5次
	for i :=0; i < numGreeters ; i++{
		go hello(&wg, i +1)

	}

	wg.Wait()


}
