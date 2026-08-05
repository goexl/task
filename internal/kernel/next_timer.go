package kernel

import (
	"time"
)

type NextTimer interface {
	Next(Task) *time.Time
}
