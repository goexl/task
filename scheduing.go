package task

import (
	"github.com/goexl/task/internal/builder"
	"github.com/goexl/task/internal/kernel"
)

func NewScheduling(target uint64, subtype kernel.Type) *builder.Scheduling {
	return builder.NewScheduling(target, subtype)
}
