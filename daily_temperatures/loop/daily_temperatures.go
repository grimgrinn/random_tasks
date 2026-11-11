package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	fmt.Println(temperatures)
	res := make([]int, len(temperatures))
	days := 0
	for i := 0; i < len(temperatures); i++ {
		for j := i + 1; j < len(temperatures); j++ {
			fmt.Printf("temperatures[i] -> %d, temperatures[j] -> %d \n", temperatures[i], temperatures[j])
			fmt.Printf("days -> %d, i -> %d, j -> %d, len(temperatures) -> %d\n", days, i, j, len(temperatures))
			if temperatures[i] < temperatures[j] {
				fmt.Printf("zapisali %d\n", days+1)
				// res = append(res, days+1)
				res[i] = j - i
				// days = 0
				break
				// } else {
				// 	days++
			}
		}
	}

	return res
}

func main() {
	fmt.Println("jopa")

	temperatures := []int{55, 38, 53, 81, 61, 93, 97, 32, 43, 78}

	res := dailyTemperatures(temperatures)

	fmt.Println(res)
}
