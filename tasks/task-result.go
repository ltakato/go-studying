package tasks

type TaskResultKey = int

const (
	PaymentProcessed TaskResultKey = iota
	TripStatusUpdated
)

type TaskResult struct {
	Key     TaskResultKey
	Message string
}
