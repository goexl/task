package kernel

import (
	"context"
	"time"
)

type Tasker interface {
	// Start 启动
	Start(ctx context.Context) error

	// Add 添加
	Add(scheduling Schedule) error

	// Remove 删除
	Remove(scheduling Schedule) error

	// Running 运行中
	Running(id uint64, status Status, retries uint32) error

	// Update 修改
	Update(id uint64, status Status, runtime time.Time) error

	// Next 更新下次执行
	Next(id uint64) error

	// Pop 取出任务并执行
	Pop(times uint32) ([]Task, bool)

	// Stop 停止
	Stop(ctx context.Context) error
}
