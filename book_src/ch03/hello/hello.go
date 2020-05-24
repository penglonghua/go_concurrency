package main

import "fmt"

// 使用 go的几种形式，这个地方都是一样的

func sayHello()  {
	fmt.Println("hello")
}

func main() {

	// 1 可以直接这样子
	go sayHello()
	// 继续执行自己的逻辑

	// 2 也可以声明一个匿名函数，并且理解执行 ，然后使用 go
	go func() {
		fmt.Println("hello")
	}()

	// 3 或者,将该函数赋值给一个变量，然后执行该变量
	sayHello2 := func() {
		fmt.Println("hello")
	}
	go sayHello2()






}
