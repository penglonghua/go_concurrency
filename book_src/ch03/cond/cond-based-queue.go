package main

// 需求说明:
// 等待信号的 go 和 发送信号的 go
// 1 假设 我们有一个固定长度为2的队列， 还有10条我们想要推送到队列中的数据.
// 我们想要在有容量(固定长度为2)的情况下，尽快地进入队列，所以就希望在队列中有空时能立即得到通知.

// 如何验证这个.

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	c := sync.NewCond(&sync.Mutex{}) // 1, 使用标准的 sync.Mutex 作为锁

	queue := make([]interface{}, 0, 10) // 10个容量的切片

	// 每隔一定时间删除 一个元素
	removeFromQueue := func(delay time.Duration) {

		time.Sleep(delay)
		c.L.Lock()
		fmt.Println("开始删除")
		queue = queue[1:] // 排除掉第一个元素, 相当与 出对
		fmt.Println("Remove from queue")
		fmt.Println("结束删除")
		c.L.Unlock()
		fmt.Println("发送信号之前")
		c.Signal() // 发送信号, 带有条件的锁，是可以发送信号的 . 我们让 一个正在等待的 goroutine 知道发生了什么事情
		fmt.Println("发送信号之后")
	}

	// 模拟10条数据
	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 { // 这个是 for循环, 当条件满足时，被锁定, 那么被锁定的是什么 !!! , 这个地方会将暂停 main goroutine ，直到一个信号的条件已经发送 .(前面的信号)
			fmt.Println("调用 wait 之前")
			c.Wait()
			fmt.Println("调用 wait 之后")
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}

}
