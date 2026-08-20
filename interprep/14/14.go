package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}

	arr2 := arr
	arr2[0] = 0
	fmt.Println(arr)
}
