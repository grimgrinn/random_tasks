package main

import "fmt"

func main() {

	nums1 := []int{4, 3, 2, 7, 8, 2, 3, 1}
	//nums := []int{1, 1}

	seen := make(map[int]bool)

	for _, n := range nums1 {
		seen[n] = true
		fmt.Println(seen)
	}

	result := []int{}

	for i := 1; i <= len(nums1); i++ {
		if !seen[i] {
			result = append(result, i)
		}
	}

	fmt.Printf("%v", result)
}
