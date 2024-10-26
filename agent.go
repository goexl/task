package task

import (
	"github.com/goexl/task/internal/builder"
	"github.com/goexl/task/internal/core"
)

// Agent 任务执行器
type Agent = core.Agent

func New(tasker core.Tasker) *builder.Agent {
	return builder.NewAgent(tasker)
}
