package main

import "sync"

func main() {
	var counter int

	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}

	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
}
