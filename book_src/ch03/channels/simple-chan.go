package main

import "fmt"

// 一个简单的 channel例子

func main() {

	stringStream := make(chan string)  //声明并定义一个 chan

	go func() {
		stringStream <- "Hello channels!" // <1>

	}()

	// 主 main go ，在接收
	fmt.Println( <- stringStream )  // <2>




}

