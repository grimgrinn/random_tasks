package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 4}

	count := make([]int, len(nums)+1)
	for _, num := range nums {
		count[num]++
	}

	twice, missing := 0, 0
	for i := 1; i <= len(nums); i++ {
		if count[i] == 2 {
			twice = i
		}
		if count[i] == 0 {
			missing = i
		}
	}

	fmt.Println([]int{twice, missing})

}
