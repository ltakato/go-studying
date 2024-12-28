package main

import (
	"context"
	"fmt"
	"go-studying/core"
	"go-studying/tasks"
)

type RunDef struct {
	Task       func(chan tasks.TaskResult)
	PostAction func()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Start processing...")

	results := make(chan tasks.TaskResult)

	sleeper := &core.RealSleeper{}
	taskRunner := tasks.New(sleeper)

	runDefs := []RunDef{
		{
			Task: taskRunner.TripStatusUpdate,
			PostAction: func() {
				fmt.Println("Processed: trip status update task")
			},
		},
		{
			Task: taskRunner.PaymentProcessing,
			PostAction: func() {
				fmt.Println("Processed: payment processing task")
				cancel()
			},
		},
	}

	for _, runDef := range runDefs {
		go func() {
			runDef.Task(results)
			runDef.PostAction()
		}()
	}

	for {
		select {
		case result := <-results:
			fmt.Println("Result:", result.Message)
		case <-ctx.Done():
			fmt.Println("Process finished")
			return
		}
	}
}
