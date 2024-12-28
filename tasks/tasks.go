package tasks

import "time"

func PaymentProcessing(results chan TaskResult) {
	time.Sleep(10 * time.Second)
	result := TaskResult{Key: PaymentProcessed, Message: "Payment processed"}
	results <- result
}

func TripStatusUpdate(results chan TaskResult) {
	time.Sleep(5 * time.Second)
	result := TaskResult{Key: TripStatusUpdated, Message: "Trip status updated"}
	results <- result
}
