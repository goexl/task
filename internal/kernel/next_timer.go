package kernel

import (
	"time"
)

type NextTimer interface {
	Next(error, Task) *time.Time
}
