package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ch <- i
		}(i)
	}
	wg.Wait()
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
}
