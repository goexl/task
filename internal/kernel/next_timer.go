package kernel

import (
	"time"
)

type NextTimer interface {
	Next(*Context, error, Task) *time.Time
}
