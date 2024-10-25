package core

import (
	"time"

	"github.com/goexl/task/internal/kernel"
)

type Tasker interface {
	// Add 添加
	Add(scheduling Scheduling) error

	// Remove 删除
	Remove(scheduling Scheduling) error

	// Running 运行中
	Running(id uint64, status kernel.Status, retries uint32) error

	// Update 修改
	Update(id uint64, status kernel.Status, runtime time.Time) error

	// Next 更新下次执行
	Next(id uint64) error
}
