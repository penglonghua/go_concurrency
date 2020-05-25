package main

// 通知其他的 go
// 信号的发送方式有之前的 Signel ，还有 Broadcast .两者不同.
// 运行时内部维护着一个 FIFO列表 (队)的方式,等待接受信号;
// Signal 发现等待最长时间的 goroutine 并通知它, 而 Broadcast 向所有的等待的 goroutine 都发送信号
//

// 例子说明.
// 我们假设正在创建一个带有按钮的 GUI应用程序.
// 我们想注册任意数量的函数，当该按钮被单击时，它将运行。
// 使用 Broadcast 方法通知所有注册的处理程序.

// 反馈
// 这个地方的顺序是不定的.
import (
	"fmt"
	"sync"
)

func main() {
	type Button struct { // <1>
		Clicked *sync.Cond
	}
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	// 2
	// 定义了一个 便利构造函数, 它允许我们注册函数处理来自条件的信号.
	// 每个处理函数都在自己的 go 上运行，并且订阅不会退出， 直到 go 被确认运行为止.
	subscribe := func(c *sync.Cond, fn func()) { // <2>
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup // <3>
	clickRegistered.Add(4)

	// 模拟最大化
	subscribe(button.Clicked, func() { // <4>
		fmt.Println("1 Maximizing window.")
		clickRegistered.Done()
	})

	// 模拟显示对话框
	subscribe(button.Clicked, func() { // <5>
		fmt.Println("2 Displaying annoying dialogue box!")
		clickRegistered.Done()
	})

	// 模拟单击鼠标
	subscribe(button.Clicked, func() { // <6>
		fmt.Println("3 Mouse clicked.")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() { // <6>
		fmt.Println("4 Mouse double  clicked.")
		clickRegistered.Done()
	})

	// 在条件锁上调用 这个 广播, 所有三个程序都将运行.
	button.Clicked.Broadcast() // <7>

	clickRegistered.Wait()
}
