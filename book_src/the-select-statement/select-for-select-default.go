package main

import (
	"fmt"
	"time"
)

// 问题
// default语句,它经常与 for-select 循环一起使用,这允许
// 一个go在等待另一个 go上报结果的同时，可以继续执行自己的操作.

// 这个地方 到底是怎么回事?
// 要到底实现什么功能那?

// 目的:
// 我们有一个循环，它在执行某种操作(自己的事情)，偶尔检查它是否应该被停止(跳出循环).

func main() {

	done := make(chan interface{})

	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0

loop:
	for {

		select {
		case <-done:
			fmt.Println("检查到某个条件成立了，不需要做事了，将退出")
			break loop // 跳出整个循环
		default: //这个地方为空的

		}

		// 模拟工作
		fmt.Println("正在做事")
		workCounter++
		time.Sleep(1 * time.Second)

	}

	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)

}
