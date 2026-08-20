package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	go func() {
		mu.Lock()
		fmt.Println("A")
		time.Sleep(2 * time.Second)
		mu.Unlock()
	}()
	go func() {
		mu.Lock()
		fmt.Println("B")
		mu.Unlock()
	}()
	time.Sleep(3 * time.Second)
}
