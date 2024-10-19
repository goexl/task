package param

import (
	"github.com/goexl/task/internal/core"
)

type Scheduling struct {
	Id      uint64
	Target  uint64
	Type    core.Type
	Subtype core.Type
	Data    any
}

func NewScheduling(target uint64, typ core.Type) *Scheduling {
	return &Scheduling{
		Target: target,
		Type:   typ,
	}
}
