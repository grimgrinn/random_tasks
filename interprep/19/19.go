package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var mu sync.Mutex
	for i := 0; i < 10; i++ {
		go func(val int) {
			mu.Lock()
			fmt.Println(i)
			mu.Unlock()
		}(i)
	}
	time.Sleep(time.Second)
}
