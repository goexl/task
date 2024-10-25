package task

import (
	"github.com/goexl/task/internal/builder"
	core2 "github.com/goexl/task/internal/internal/core"
	"github.com/goexl/task/internal/kernel"
)

// Scheduling 计划
type Scheduling = core2.Scheduling

func NewScheduling(target uint64, subtype kernel.Type) *builder.Scheduling {
	return builder.NewScheduling(target, subtype)
}
