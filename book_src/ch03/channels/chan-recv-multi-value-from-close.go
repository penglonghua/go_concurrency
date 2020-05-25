package main

import "fmt"

// chan 中返回2个值
// 从一个 关闭的 chan，也能读取,原因在于，有其默认的零值

func main() {

	intStream := make(chan int)
	close(intStream)



	salutation, ok := <-intStream
	fmt.Printf("(%v): %v\n", ok, salutation)

}