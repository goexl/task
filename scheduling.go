package task

import (
	"github.com/goexl/task/internal/builder"
)

type Scheduling interface {
	// Id 任务标识，如果不提供将自动生成
	Id() uint64

	// Target 目标
	Target() uint64

	// Data 数据
	Data() any
}

func NewScheduling(target uint64) *builder.Scheduling {
	return builder.NewScheduling(target)
}
