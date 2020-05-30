package main

import (
	"fmt"
	"time"
)

// 这里演示一个 防止 go泄漏的例子.

// 将 父子go 进行成功整合的一种方法就是在
// 父子 go之间建立一个 "信号通道", 让 父go 向子 go发出取消信号.
// 按照惯例， 这个信号通常是一个名叫 done 的只读 channel .
// 父 go将该 channel传递给 子go, 然后在想要取消子go时，关闭该 channel.

func main() {
	doWork := func(
		done <-chan interface{},
		strings <-chan string,
	) <-chan interface{} { // <1>
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					// Do something interesting
					fmt.Println(s)
				case <-done: // <2>
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() { // <3>
		// Cancel the operation after 1 second.
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	<-terminated // <4>
	fmt.Println("Done.")
}
