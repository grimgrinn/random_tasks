package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 5, 1, 3, 4, 7}
	n := 3

	res := []int{}

	for i, j := 0, n; i < n; i, j = i+1, j+1 {
		fmt.Println(i, j)
		res = append(res, nums[i], nums[j])
		fmt.Println(res)
	}

	fmt.Println(res)

}
