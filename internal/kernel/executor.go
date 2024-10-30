package kernel

import (
	"github.com/goexl/task/internal/internal/core"
)

type Executor interface {
	// Execute 执行任务
	Execute(ctx *core.Context, target uint64, retries uint32) error

	// Next 更新下次执行
	Next(id uint64) error

	// Error 执行错误
	Error() error
}
