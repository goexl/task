package param

import (
	"time"

	"github.com/goexl/task/internal/kernel"
)

type Schedule struct {
	Id      uint64
	Target  uint64
	Type    kernel.Type
	Subtype kernel.Type
	Timeout time.Duration
	Maximum uint64
	Data    map[string]any
}

func NewSchedule(target uint64, subtype kernel.Type) *Schedule {
	return &Schedule{
		Target:  target,
		Subtype: subtype,
		Timeout: 10 * time.Minute, // 默认十分钟必须执行完成
		Data:    make(map[string]any),
	}
}
