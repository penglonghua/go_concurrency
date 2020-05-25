package main

import (
	"fmt"
	"sync"
)

// 问题， 下面的输出是什么

// 注意
// sync.Once 只计算调用 Do方法的次数,而不是多少次调用Do中的传递的方法
// 这样子， sync.Once的副本与所调用的函数紧密耦合
func main() {
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once

	once.Do(decrement)
	once.Do(increment)

	fmt.Printf("Count: %d\n", count)
}
