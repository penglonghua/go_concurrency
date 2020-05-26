package main

import "fmt"

// 尽量保持 channel所有权的范围很小
// 这个程序也是一个非常好的例子

func main() {

	// 这里声明一个变量，函数类型，返回的是 消费者
	chanOwner := func() <- chan int {

		resultStream := make(chan int, 5) // 1
		go func() {      // 2
			defer close(resultStream)  // 3
			for i := 0; i <=5; i++{
				resultStream <- i    // 发送
			}
		}()

		return resultStream // 4 这个地方会隐士的转换成 消费者


	}


	resultStream := chanOwner()

	// 消费者
	for result := range resultStream{ // 5

		fmt.Printf("Received: %d\n", result)

	}
	fmt.Println("Done receiving!")

}