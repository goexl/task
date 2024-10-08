package task

import (
	"context"

	"github.com/goexl/task/internal/core"
)

// Agent 任务执行器
type Agent interface {
	// Start 开始接收任务
	Start(ctx context.Context, processor Processor) error

	// Add 添加
	Add(scheduling core.Scheduling) error

	// Remove 删除
	Remove(scheduling core.Scheduling) error

	// Stop 停止
	Stop(ctx context.Context) error
}
