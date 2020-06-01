package main

import "fmt"

// 构建 pipeline的简单例子, 第一个
// 这个地方作为对比，这个地方的 stage 是批处理，对大块数据进行一次操作，而不是一次一个值.

func main() {

	// 这里构建一个 stage
	multiply := func(values []int, multiplier int) []int {

		multipliedValues := make([]int, len(values))

		for i, v := range values {
			multipliedValues[i] = v * multiplier
		}

		return multipliedValues

	}

	// 这里在构建另一个 stage
	add := func(values []int, additive int) []int {
		addedValues := make([]int, len(values))
		for i, v := range values {
			addedValues[i] = v + additive
		}
		return addedValues
	}

	// 现在组合这两个 stage
	ints := []int{1, 2, 3, 4}
	for _, v := range multiply(add(multiply(ints, 2), 1), 2) {
		fmt.Println(v)
	}

}
