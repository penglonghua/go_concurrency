package main

import "fmt"

// 程序如何与 值为 nil的 channel交互?

func main() {

	// 1 尝试从 nil channel 中读取数据

	var dataStream  chan interface{}   // 这个地方只有声明，并没有定义,所以是 nil
	// var name type  就是只有声明，没有定义

	fmt.Printf("%T", dataStream)

	//<-dataStream  // 读取数据
	//  all goroutines are asleep - deadlock!

	// 2 那么写入
	//dataStream <- "xxx"
	// all goroutines are asleep - deadlock!

	// 3
	close(dataStream)



}