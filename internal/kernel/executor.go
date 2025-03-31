package kernel

import (
	"github.com/goexl/id"
)

type Executor interface {
	// Execute 执行任务
	Execute(ctx *Context, target id.Value, times uint32) error
}
