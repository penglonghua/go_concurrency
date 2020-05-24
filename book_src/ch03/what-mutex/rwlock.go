package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

//  先说下原因
// 进入和退出一个临界区是有消耗的，所以一般人都会尽量减少在临界区的时间.
// 这个地方 读写锁还有其他更好的例子

func main() {

	//	生成者

	producer := func(wg *sync.WaitGroup, l sync.Locker) {

		defer wg.Done()

		for i := 5; i >= 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1) // 让这个生成者低效一点
		}

	}

	// 观察者
	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	//测试,返回的时间

	test := func(count int, mutex sync.Locker, rwMutex sync.Locker) time.Duration {

		var wg sync.WaitGroup
		wg.Add(count + 1)
		beginTestTime := time.Now()
		go producer(&wg, mutex)
		for i := count; i > 0; i-- {
			go observer(&wg, rwMutex)
		}

		wg.Wait()
		return time.Since(beginTestTime)
	}

	// 输出到表格数据
	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutext\tMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(tw, "%d\t%v\t%v\n", count, test(count, &m, m.RLocker()), test(count, &m, &m))
	}

}
