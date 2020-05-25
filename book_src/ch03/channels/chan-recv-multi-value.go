package main

import "fmt"

// chan 中返回2个值

func main() {

	stringStream := make(chan string)

	go func() {
		stringStream <- "hello channels!"
	}()

	salutation, ok := <- stringStream
	fmt.Printf("(%v): %v\n", ok, salutation)

}