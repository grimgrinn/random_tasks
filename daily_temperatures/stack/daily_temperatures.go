package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	res := make([]int, len(temperatures))

	for i, temp := range temperatures {
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			res[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return res
}
func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	res := dailyTemperatures(temperatures)

	fmt.Println(res)
}
