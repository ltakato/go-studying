package tasks

import (
	"go-studying/core"
	"time"
)

type TaskRunner struct {
	sleeper core.Sleeper
}

func New(sleeper core.Sleeper) *TaskRunner {
	taskRunner := TaskRunner{
		sleeper: sleeper,
	}
	return &taskRunner
}

func (t *TaskRunner) PaymentProcessing(results chan TaskResult) {
	t.sleeper.Sleep(10 * time.Second)
	result := TaskResult{Key: PaymentProcessed, Message: "Payment processed"}
	results <- result
}

func (t *TaskRunner) TripStatusUpdate(results chan TaskResult) {
	t.sleeper.Sleep(5 * time.Second)
	result := TaskResult{Key: TripStatusUpdated, Message: "Trip status updated"}
	results <- result
}
