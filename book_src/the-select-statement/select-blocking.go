package main

import (
	"fmt"
	"time"
)

// 这个地方是最简单的 select 语句,看看有什么作用

func main() {

	start := time.Now()
	c := make(chan interface{})

	go func() {
		time.Sleep(5 * time.Second)
		close(c) // <1>
	}()

	fmt.Println("Blocking on read...")

	// 进入到 select 后就等待了,然后有 case 执行，然后就退出了
	select {
	case <-c: // <2>
		fmt.Printf("Unblocked %v later.\n", time.Since(start))

	}

	//<-c // 也可以直接用这个.

	//fmt.Printf("Unblocked %v later.\n", time.Since(start))

}
