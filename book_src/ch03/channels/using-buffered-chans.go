package main

import (
	"bytes"
	"fmt"
	"os"
)

// 使用带缓冲的 chan
// 这个地方的例子可能不太恰当，但是可以帮忙了解一点点

func main() {

	var stdoutBuff bytes.Buffer //<1> 带缓冲的内存
	defer stdoutBuff.WriteTo(os.Stdout) //<2>

	intStream := make(chan int, 4) // <3>
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")

		for i := 0; i < 5; i++{
			fmt.Fprintf(&stdoutBuff, "Sendint : %d\n", i)
			intStream <- i
		}
	}()




	// 遍历 channel
	for integer := range intStream{
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)

	}





}