// the or-channel
// 需求是什么:
// 有时候，你可能发现你想要将多个 done channel 合并成一个单一的 done channel
// 来这个复合 done channel 中的任意一个 channel 关闭的时候,关闭整个 done channel.
package main

import (
	"fmt"
	"time"
)

func main() {
	// 这种模式通过递归和 goroutine 创建一个复合 done channel

	// 声明
	var or func(channels ...<-chan interface{}) <-chan interface{}

	//定义
	or = func(channels ...<-chan interface{}) <-chan interface{} {

		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]

		}

		// 返回值
		orDone := make(chan interface{})

		// 这是函数的主体，以及递归发生的地方.
		//我们启动一个新的 go来让我们的 channel 可以不被阻塞地等待消息的到来
		go func() {
			defer close(orDone)

			switch len(channels) {

			case 2: // <5>
				select {
				case <-channels[0]:
				case <-channels[1]:
				}

			// <6>
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...): // <6> append 返回的是切片，然后在进行 解包
				}

			}

		}()

		return orDone
	}

	// 下面开始 调用了
	// 上面是一个相当简洁的函数，使你可以将任意数量的 channel组合到的单个 channel
	// 只要任何组件中 channel关闭或写入(没有测试过),该 channel就会关闭.

	// 它将经过一段时间后关闭的 channel, 并将这些 channel合并到一个关闭的单个 chanel中:
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now() // <2>
	//测试
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v\n", time.Since(start)) // <3>

}
