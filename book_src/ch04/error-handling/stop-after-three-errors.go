package main

import (
	"fmt"
	"net/http"
)

// 在前面的例子中，我们只是将错误写入 stdio,但是我们可以做其他的事情。
// 让我们稍微修改我们的程序，以便在出现三个或更多错误时停止尝试检查状态:
func main() {
	type Result struct { // <1>
		Error    error
		Response *http.Response
	}
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result { // <2>
		results := make(chan Result)
		go func() {
			defer close(results)

			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp} // <3>
				select {
				case <-done:
					return
				case results <- result: // <4>
				}
			}
		}()
		return results
	}
	done := make(chan interface{})
	defer close(done)

	// 上面还是之前的代码，只是多个，检查次数
	errCount := 0
	urls := []string{"a", "https://www.google.com", "b", "c", "d"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
			errCount++
			if errCount >= 3 {
				fmt.Println("Too many errors, breaking!")
				break
			}
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
