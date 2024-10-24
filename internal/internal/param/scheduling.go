package param

import (
	"github.com/goexl/task/internal/core"
)

type Scheduling struct {
	Id      uint64
	Target  uint64
	Type    core.Type
	Subtype core.Type
	Data    map[string]any
}

func NewScheduling(target uint64, subtype core.Type) *Scheduling {
	return &Scheduling{
		Target:  target,
		Subtype: subtype,
		Data:    make(map[string]any),
	}
}
