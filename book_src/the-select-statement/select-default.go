package main

import (
	"fmt"
	"time"
)

// select 中的 default 语句
// 当 所有的 case语句都不满足时，就是选择 default语句.
// 但是因为，又是并发的，又是 随机的，所以这个地方 的  default几乎一瞬间执行了.

// 那现在问题是什么
// 因为有 default语句，它几乎是瞬间运行了默认语句。这允许在不阻塞的情况下退出 select模块.

// 通常情况下，它与  for -select 一块循环使用, 后面的例子
//其他的当然也有机会使用，但是还是一瞬间的事情, 太不可控
func main() {

	start := time.Now()
	var c1, c2 <-chan int

	select {

	case <-c1:

	case <-c2:
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))

	}

}
