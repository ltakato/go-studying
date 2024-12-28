package main

import (
	"context"
	"fmt"
	"go-studying/tasks"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("start processing...")

	results := make(chan tasks.TaskResult)

	tasksDefs := []func(chan tasks.TaskResult){
		tasks.TripStatusUpdate,
		tasks.PaymentProcessing,
	}

	for _, task := range tasksDefs {
		go task(results)
	}

	for {
		select {
		case result := <-results:
			fmt.Println("Result:", result.Message)
			if result.Key == tasks.PaymentProcessed {
				cancel()
			}
		case <-ctx.Done():
			fmt.Println("Process finished")
			return
		}
	}
}
