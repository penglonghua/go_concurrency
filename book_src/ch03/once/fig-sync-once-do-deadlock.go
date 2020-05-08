package main

import (
	"sync"
)

// all goroutines are asleep - deadlock!

func main() {
	var onceA, onceB sync.Once
	var initB func() // 这是一个函数类型
	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) } // <1>
	onceA.Do(initA)                    // <2>
}
