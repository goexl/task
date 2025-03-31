package task

import (
	"github.com/goexl/id"
	"github.com/goexl/task/internal/builder"
	"github.com/goexl/task/internal/kernel"
)

// Schedule 计划
type Schedule = kernel.Schedule

func NewSchedule(target id.Value, subtype kernel.Type) *builder.Schedule {
	return builder.NewSchedule(target, subtype)
}
