package main

import (
	"context"
	"fmt"
	"go-studying/core"
	"go-studying/tasks"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Start processing...")

	results := make(chan tasks.TaskResult)

	sleeper := &core.RealSleeper{}
	taskRunner := tasks.New(sleeper)

	tasksDefs := []func(chan tasks.TaskResult){
		taskRunner.TripStatusUpdate,
		taskRunner.PaymentProcessing,
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
