package main

import (
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// go func() {
		// 	defer wg.Done()
		// 	panic("oops")
		// }()
		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Recover: %v", r)
				}
				wg.Done()
			}()
			panic("oops")
		}()
		wg.Wait()
	}

}
