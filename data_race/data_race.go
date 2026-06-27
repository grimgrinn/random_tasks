package main

import (
	"fmt"
	"sync"
)

func main() {
	var money int32
	var donationCount int32

	mutex := &sync.RWMutex{}

	wg := &sync.WaitGroup{}

	go func() {
		for {
			mutex.RLock()
			m := money
			dc := donationCount
			mutex.RUnlock()

			if m != dc {
				fmt.Println("money=", m, "donations=", dc)
				break
			}
		}
	}()

	wg.Add(3000)
	for range 3000 {
		go func() {
			defer wg.Done()

			mutex.Lock()
			money++
			donationCount++
			mutex.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(money)
}
