package main

import "fmt"

// 多个 channel 同时可用时，select 会选择谁, 它的策略是什么

// 如何做实验
// 测试结果， 基本上是一半一半
func main() {

	c1 := make(chan interface{})
	close(c1)

	c2 := make(chan interface{})
	close(c2)

	var c1Count, c2Count int
	for i := 1000; i >= 0; i-- {
		// 测试1000次，看看 select 的选择会是什么
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}

	}
	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)

}
