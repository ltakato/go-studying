package tasks

import (
	"fmt"
	"testing"
	"time"
)

type MockSleeper struct{}

func (ms *MockSleeper) Sleep(duration time.Duration) {
	fmt.Printf("Mock sleep called for %v\n", duration)
}

func TestNew(t *testing.T) {
	t.Run("should instance a task runner", func(t *testing.T) {
		sleeper := &MockSleeper{}
		taskRunner := New(sleeper)

		if taskRunner.sleeper == nil {
			t.Errorf("expected to instance with sleeper")
		}
	})
}

func TestTripStatusUpdate(t *testing.T) {
	t.Run("should run trip status update task", func(t *testing.T) {
		sleeper := &MockSleeper{}
		taskRunner := New(sleeper)
		results := make(chan TaskResult)
		go taskRunner.TripStatusUpdate(results)

		select {
		case result := <-results:
			expected := TripStatusUpdated
			if result.Key != expected {
				t.Errorf("Expected %d, got %d", expected, result.Key)
			}
		case <-time.After(3 * time.Second):
			t.Fatal("Test timed out")
		}
	})
}

func TestPaymentProcessing(t *testing.T) {
	t.Run("should run payment processing task", func(t *testing.T) {
		sleeper := &MockSleeper{}
		taskRunner := New(sleeper)
		results := make(chan TaskResult)
		go taskRunner.PaymentProcessing(results)

		select {
		case result := <-results:
			expected := PaymentProcessed
			if result.Key != expected {
				t.Errorf("Expected %d, got %d", expected, result.Key)
			}
		case <-time.After(3 * time.Second):
			t.Fatal("Test timed out")
		}
	})
}
