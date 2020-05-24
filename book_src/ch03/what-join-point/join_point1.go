package main

import (
	"fmt"
	"sync"
)

// 其中一种 实现 join point 的方式,应该还会其他的方式

func main() {


	// 1 定义一个 等待组
	var wg sync.WaitGroup

	sayHello := func() {
		//2  wg.Done() // 在实际的执行 go 里面，表示这个 wg 做完了 , -1
		fmt.Println("hello")
	}

	// 3 在将要执行 go 之前，使 go + 1
	wg.Add(1)

	// 开始 一个 go
	go sayHello()

	// 4 这个地方就是 join point, 会阻塞等待 直到 goroutine 托管 sayHello函数为止
	wg.Wait()


}




