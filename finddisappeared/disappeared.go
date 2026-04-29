package main

import "fmt"

func main() {

	// nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	nums := []int{1, 1}

	arr := make([]int, len(nums))

	ans := []int{}
	fmt.Println(len(nums))
	fmt.Println(arr)
	fmt.Println(nums)

	isin := make(map[int]int)

	for i, n := range nums {

		if _, exists := isin[n]; !exists {
			isin[n] = i

			// fmt.Println(isin)
		} else {
			fmt.Println(i, "!")
			ans = append(ans, i)

		}
	}

	fmt.Println(isin)

	fmt.Println(ans)

}
