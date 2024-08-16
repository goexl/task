package task

import (
	"github.com/goexl/task/internal"
)

type Scheduling interface {
	// Target 目标
	Target() int64

	// Data 数据
	Data() any
}

func NewScheduling(target int64, data any) Scheduling {
	return internal.NewScheduling(target, data)
}
