package main

import "fmt"

var nums1 = []int{4, 3, 2, 7, 8, 2, 3, 1}
var nums2 = []int{1, 1}

// O(n log n) + O(n)
func findDisappearedNumbers1(nums []int) []int {
	seen := make(map[int]bool)

	for _, n := range nums {
		seen[n] = true
	}

	result := []int{}

	for i := 1; i <= len(nums); i++ {
		if !seen[i] {
			result = append(result, i)
		}
	}

	return result
}

// O(n), O(1)
func findDisappearedNumbers2(nums []int) []int {
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		index := abs(nums[i]) - 1
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}

		fmt.Println(nums)
	}

	result := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}

	return result
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// fmt.Println(findDisappearedNumbers1(nums1))
	// fmt.Println(findDisappearedNumbers1(nums2))

	fmt.Println(findDisappearedNumbers2(nums1))

}
