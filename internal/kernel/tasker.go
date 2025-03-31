package kernel

import (
	"context"
	"time"

	"github.com/goexl/id"
)

type Tasker interface {
	// Start 启动
	Start(ctx context.Context) error

	// Add 添加
	Add(schedule Schedule, schedules ...Schedule) error

	// Remove 删除
	Remove(schedule Schedule) error

	// Running 运行中
	Running(id id.Value, status Status, times uint32) error

	// Update 修改
	Update(id id.Value, status Status, runtime time.Time) error

	// Pop 取出任务并执行
	Pop() Task

	// Archive 存档
	Archive(task Task) error

	// Failed 执行错误
	Failed(task Task) error

	// Stop 停止
	Stop(ctx context.Context) error
}
