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

	// 用 4K初始化 pool
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

			mem := calcPool.Get().(*[]byte) // 2 FIXME 断言类型
			defer calcPool.Put(mem)         //放回去

			// Assume something interesting, but quick is being done with
			// this memory.

		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created.", numCalcsCreated)

}
