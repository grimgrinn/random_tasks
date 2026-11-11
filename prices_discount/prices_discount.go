package main

import "fmt"

func finalPrices(prices []int) []int {
	res := []int{}

	for i := 0; i < len(prices); i++ {
		discountFound := false

		for j := i + 1; j < len(prices); j++ {
			if prices[i] >= prices[j] {
				res = append(res, prices[i]-prices[j])
				discountFound = true
				break
			}

		}
		if !discountFound {
			res = append(res, prices[i])
		}
	}

	fmt.Println(prices)
	return res
}

func main() {
	prices := []int{8, 4, 6, 2, 3}

	res := finalPrices(prices)
	fmt.Println(res)
}
