package main

// 这里是一个最简单的例子来说明:

import (
	"fmt"
	"sync"
)

func main() {

	// myPool 是一个指针
	// 里面有一个 new 函数
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	myPool.Get()             // <1>  //如果有实例,那么直接返回该实例，如果没有实例， 调用 pool中的 new函数
	instance := myPool.Get() //<2>
	myPool.Put(instance)     //<3>
	myPool.Get()             //<4>
	myPool.Get()
	myPool.Get()

}
