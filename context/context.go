package main

import "context"

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return // Завершаемся
		default:
			// работаем
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)

	// когда нужно остановить
	cancel()
}
