package param

import (
	"github.com/goexl/task/internal/core"
)

type Scheduling struct {
	Id      uint64
	Target  uint64
	Type    core.Subtype
	Subtype core.Subtype
	Data    any
}

func NewScheduling(target uint64, typ core.Subtype) *Scheduling {
	return &Scheduling{
		Target: target,
		Type:   typ,
	}
}
