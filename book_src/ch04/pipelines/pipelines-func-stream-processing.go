package main

import "fmt"

// 这个地方作为对象, 使用 pipeline 执行流处理(也就是一次只接收和处理一个元素.)
// 而不是 pipeline 的另一种，批处理

func main() {

	multiply := func(value, multiplier int) int {

		return value * multiplier

	}

	add := func(value, additive int) int {
		return value + additive
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range ints {
		fmt.Println(multiply(add(multiply(v, 2), 1), 2))
	}

}
