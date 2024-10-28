package task

import (
	"github.com/goexl/task/internal/builder"
	"github.com/goexl/task/internal/core"
	"github.com/goexl/task/internal/kernel"
)

// Scheduling 计划
type Scheduling = core.Scheduling

func NewScheduling(target uint64, subtype kernel.Type) *builder.Scheduling {
	return builder.NewScheduling(target, subtype)
}
