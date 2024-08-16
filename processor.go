package task

import (
	"time"
)

// Processor 处理器
type Processor interface {
	// Process 处理任务调度
	Process(scheduling Scheduling) error

	// Recyclable 是否继续执行
	Recyclable() bool

	// Next 下次执行时间
	Next() time.Time

	// Error 执行错误
	Error() error
}
