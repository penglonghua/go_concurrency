package main

import (
	"fmt"
	"sync"
)

func main() {

	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			return &mem // <1>  返回的是 存储 bytes 切片的地址
		},
	}

	// 用 4K初始化 pool, 4个,每个都是 1K
	// 这里初始化了 4次
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()

			//这个地方也会再次 4次, 类型断言这个地方,因为 前面就有 4次
			mem := calcPool.Get().(*[]byte) // 2 FIXME 断言类型, 语法如下去 obj.(type) 返回值是什么?
			defer calcPool.Put(mem)         //放回去 , 而且这个地方是 defer 做的事情

			// Assume something interesting, but quick is being done with
			// this memory.

		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created.\n", numCalcsCreated)

}
