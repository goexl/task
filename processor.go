package task

import (
	"github.com/goexl/task/internal/core"
)

// Processor 处理器
type Processor interface {
	// Process 处理任务调度
	Process(core.Tasking) (Executor, error)
}
