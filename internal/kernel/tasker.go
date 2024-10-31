package kernel

import (
	"context"
	"time"
)

type Tasker interface {
	// Start 启动
	Start(ctx context.Context) error

	// Add 添加
	Add(schedule Schedule) error

	// Remove 删除
	Remove(schedule Schedule) error

	// Running 运行中
	Running(id uint64, status Status, retries uint32) error

	// Update 修改
	Update(id uint64, status Status, runtime time.Time) error

	// Pop 取出任务并执行
	Pop(retries uint32) []Task

	// Archive 存档
	Archive(task Task) error

	// Faield 执行错误
	Faield(task Task) error

	// Stop 停止
	Stop(ctx context.Context) error
}
