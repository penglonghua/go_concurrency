package main

import (
	"fmt"
	"sync"
)

// 什么叫锁
// 临界区的概念, 就是程序中需要独占共享资源的区域

func main() {

	var count int       // 共享变量
	var lock sync.Mutex //互斥锁

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing:%d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing:%d\n", count)
	}

	//	开始并发了

	var wg sync.WaitGroup

	// 增加
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()

		}()
	}

	// 减少
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decrement()
		}()
	}

	wg.Wait()
	fmt.Println("Arithmetic complete")

}
