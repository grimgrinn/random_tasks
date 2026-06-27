package main

import (
	"fmt"
	"time"
)

// Написать 3 функции:
// writer - генерит числа от 1 до 10
// doubler - умножает числа на 2, имитируя работу (500ms)
// reader - читает и выводит на экран

func writer() <-chan int {
	ch := make(chan int)

	go func() {
		for i := range 10 {
			ch <- i + 1
		}
		close(ch)
	}()

	return ch
}

func doubler(inputCh <-chan int) <-chan int {
	db := make(chan int)

	go func() {
		for n := range inputCh {
			time.Sleep(500 * time.Millisecond)
			db <- n * 2
		}
		close(db)
	}()

	return db
}

func reader(ch <-chan int) {
	for n := range ch {
		fmt.Println(n)
	}
}
func main() {
	reader(doubler(writer()))
}
