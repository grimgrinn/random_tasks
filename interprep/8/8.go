package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1) // добавляем до запуска
		go func() {
			//		wg.Add(1)
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}
