package main

import (
	"fmt"
	"net/http"
)

// 正确的处理方式
// checkStatus 中产生的所有的结果集都可以被传递给 我们的 "main go",来对各种可能的错误进行处理
// 从更广泛的角度来说, 我们已经将错误处理的担忧从我们的生产者 goroutine 中分离出来了。

func main() {

	type Result struct {
		Error    error
		Response *http.Response
	}

	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)

		go func() {
			defer close(results)

			for _, url := range urls {

				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp}

				select {
				case <-done:
					return
				case results <- result: //<4> // 写,发送语句
				}

			}

		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"http://www.baidu.com", "https://www.google.com", "https://badhost"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}

}
