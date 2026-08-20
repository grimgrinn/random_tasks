package main

import "fmt"

func main() {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go func(i int) {
			ch <- i
		}(i)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}
