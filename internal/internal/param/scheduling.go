package param

import (
	"github.com/goexl/task/internal/kernel"
)

type Scheduling struct {
	Id        uint64
	Target    uint64
	Type      kernel.Type
	Subtype   kernel.Type
	Removable bool
	Data      map[string]any
}

func NewScheduling(target uint64, subtype kernel.Type) *Scheduling {
	return &Scheduling{
		Target:    target,
		Subtype:   subtype,
		Removable: false,
		Data:      make(map[string]any),
	}
}
