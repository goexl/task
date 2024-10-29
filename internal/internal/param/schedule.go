package param

import (
	"time"

	"github.com/goexl/task/internal/kernel"
)

type Schedule struct {
	Id      uint64
	Target  uint64
	Type    kernel.Type
	Subtype kernel.Type
	Elapsed time.Duration
	Data    map[string]any
}

func NewSchedule(target uint64, subtype kernel.Type) *Schedule {
	return &Schedule{
		Target:  target,
		Subtype: subtype,
		Elapsed: 24 * time.Hour,
		Data:    make(map[string]any),
	}
}
