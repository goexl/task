package core

import (
	"github.com/goexl/task/internal/kernel"
)

type Scheduling interface {
	// Id 任务标识，如果不提供将自动生成
	Id() uint64

	// Target 目标
	Target() uint64

	// Removable 是否可被删除
	Removable() bool

	// Type 类型
	Type() kernel.Type

	// Subtype 子类型
	Subtype() kernel.Type

	// Data 数据
	Data() map[string]any
}
