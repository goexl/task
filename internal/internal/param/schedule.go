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
	Timeout time.Duration
	Maximum uint32
	Data    map[string]any
}

func NewSchedule(target uint64, subtype kernel.Type) *Schedule {
	return &Schedule{
		Target:  target,
		Subtype: subtype,
		Timeout: 24 * time.Hour,
		Maximum: 10,
		Data:    make(map[string]any),
	}
}
