package main

import "fmt"

func largestRectangleArea(heights []int) int {
	fmt.Println(heights)

	heights = append(heights, 0)
	stack := []int{}
	maxArea := 0

	for i, h := range heights {
		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			heightIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			height := heights[heightIndex]

			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}
	return maxArea
}

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}

	fmt.Println(largestRectangleArea(heights))
}
