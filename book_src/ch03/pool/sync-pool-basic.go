package main

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

	myPool.Get()             // <1>
	instance := myPool.Get() //<2>
	myPool.Put(instance)     //<3>
	myPool.Get()             //<4>

}
