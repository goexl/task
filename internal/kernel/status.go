package kernel

const (
	StatusCreated  = 1
	StatusRunning  = 2
	StatusRetrying = 3

	StatusFailed = 10

	StatusSuccess = 20
)

type Status uint8
