package main

// 前面演示的是 消费者，如果换成是生产者，还有什么问题.
// 问题：
//3 random ints:
//1: 5577006791947779410
//2: 8674665223082153551
//3: 6129484611666145821
// 发现，没有输出 newRandStream closure exited.

import (
	"fmt"
	"math/rand"
)

func main() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.") // <1>
			defer close(randStream)
			for {
				randStream <- rand.Int() // 发送语句
			}
		}()

		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
}
