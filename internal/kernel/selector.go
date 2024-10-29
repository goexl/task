package kernel

import (
	"github.com/goexl/task/internal/core"
)

type Selector interface {
	// Select 选择不同的执行器执行任务
	Select(core.Task) (core.Executor, error)
}
