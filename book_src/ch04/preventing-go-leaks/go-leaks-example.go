package main

import "fmt"

// 这个地方演示一个 go泄漏的例子

// 虽然程序也可以运行, 但是这个地方会泄漏
func main() {

	doWork := func(strings <-chan string) <-chan interface{} {

		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				// Do something interesting
				fmt.Println(s)
			}
		}()
		return completed

	}

	doWork(nil) // 这个地方传入的是 nil.

	// Perhaps more work is done here
	fmt.Println("Done.")

}
