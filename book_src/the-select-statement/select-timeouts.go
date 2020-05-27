package main

import (
	"fmt"
	"time"
)

// select 中如何没有任何可用的 channle ,会发生阻塞， 如果不想被永远阻塞下去
// 可以使用 超时机制，go语言中 Time包可以优雅的提供

func main() {

	var c <-chan int // 接受,消费

	start := time.Now()
	select {
	case <-c: // <1> 这个 case 永远不会被执行，因为我们是从 nil channel 读取的

	// time.After 的返回值是 <- chan time, 消费者
	case <-time.After(3 * time.Second):
		fmt.Printf("Time out %v later.\n", time.Since(start))

	}

}
