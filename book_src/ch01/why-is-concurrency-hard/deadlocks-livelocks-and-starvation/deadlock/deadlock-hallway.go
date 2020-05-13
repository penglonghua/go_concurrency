/**
* 什么叫死锁
*  并发进程彼此等待, 这是一种死锁.
*  下面模拟这个情况
* 输出 fatal error: all goroutines are asleep - deadlock!
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu sync.Mutex
	value int
}




func main() {

	var wg sync.WaitGroup

	printSum := func(v1, v2 *value) {
		defer  wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 *time.Second) // <3> 休眠一段时间来模拟一些工作（并触发死锁）.

		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value + v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)  // a, b
	go printSum(&b, &a)  // b, a 注意这个地方反起来的
	wg.Wait()


}





