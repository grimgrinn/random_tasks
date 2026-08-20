package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stopped")
				return
			default:
				fmt.Println("working")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(500 * time.Millisecond)
	cancel()
	time.Sleep(time.Second)
}
