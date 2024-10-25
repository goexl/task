package task

import (
	"context"

	core2 "github.com/goexl/task/internal/internal/core"
)

// Agent 任务执行器
type Agent interface {
	// Start 开始接收任务
	Start(ctx context.Context, processor core2.Processor) error

	// Add 添加
	Add(scheduling core2.Scheduling) error

	// Remove 删除
	Remove(scheduling core2.Scheduling) error

	// Stop 停止
	Stop(ctx context.Context) error
}
