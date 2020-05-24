package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 讨论一个关于 goroutine 有趣的事情: GC并没有回收被丢弃的 goroutine.
//
//这里有个例子:
//我们将利用 goroutine 不被GC的事实与运行时的自省能力结合起来，并测算
//goroutine 的大小.

//

func main() {

	// 内存使用
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	// 定义一个 接收 chan
	var c <- chan interface{}

	var wg sync.WaitGroup
	noop := func() {wg.Done(); <-c}
	// 这个 go ，因为里面有接收语句，而实际上没有发送，所有
	// 会导致阻塞，所以会 无限等待

	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memConsumed()

	for i := numGoroutines; i > 0; i--{
		go noop()
	}

	wg.Wait()

	after := memConsumed()
	// 计算一个 go 占用了多少内存

	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
	// 输出 0.059kb
	//输出的 0.033kb
	// 发现，很小，而且不是固定的
}