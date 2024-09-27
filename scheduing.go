package task

import (
	"github.com/goexl/task/internal/builder"
	"github.com/goexl/task/internal/core"
)

// Scheduling 计划
type Scheduling = core.Scheduling

func NewScheduling(target uint64) *builder.Scheduling {
	return builder.NewScheduling(target)
}
