package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	s2 := s
	s2 = append(s2, 4)
	fmt.Println(s)
	fmt.Println(s2)

}
