package main

import "fmt"

func main() {

	intStream := make(chan int)

	go func() {
		defer  close(intStream) // 确保 go在退出之前 channel是关闭的.
		// 如果不关闭，会打印死锁
		for i := 1; i <= 5; i++{
			intStream <- i
		}
	}()


	// 遍历
	// 注意  range 这个地方不返回 第2个 布尔值
	for integer := range intStream{
		fmt.Printf("%v ", integer)
	}

}
