package main

import "fmt"

// 我们可以从一个 已经关闭的 channel读取数据，
// 这里提出一个问题
// close 肯定可以关闭 上游的 生产者, 问题是 能否关闭消费者

func main() {

	intStream := make(chan int)
	close(intStream)

	salutation, ok := <-intStream // <1>
	fmt.Printf("(%v): %v\n", ok, salutation)

	// 这个地方 观察发现, close 函数的签名 ,就是一个 发送语句, 猜测不能 关闭消费者
	// func close(c chan<- Type)
	// 如果消费者有多个？关闭还有意义码？

	readStream := make(<- chan interface{} ) //消费者
	close(readStream)
	// # command-line-arguments
	//./chan-close.go:22:7: invalid operation: close(readStream) (cannot close receive-only channel)

	// 答案已经有了，不能关闭一个 接收的(消费者)channel.






}

