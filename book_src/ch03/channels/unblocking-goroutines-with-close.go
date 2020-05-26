package main

import (
	"fmt"
	"sync"
)

// 这个地方得慢慢理一下,重点是理解要先到位

// 关闭 channel 也是一种同时给多个 goroutine 发信号的方法.

// 由于一个被关闭的 channel可以被无数次读取，所以不管有多少 goroutine 在等待它.
// 由于调用了 close,这个地方也相当于 同一时间打开了多个 等待的 go .

func main() {

	begin := make(chan interface{})

	var wg sync.WaitGroup
	for i := 0; i < 5; i++{
		wg.Add(1)
		go func(i int) {
			defer  wg.Done()
			<- begin // <1> 接收
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	// 先打印出 Unblocking goroutines...
	fmt.Println("Unblocking goroutines...")
	close(begin)  // <2>  // 这个调用之后，才能  <- begin 那个才会不阻塞
	wg.Wait()







}