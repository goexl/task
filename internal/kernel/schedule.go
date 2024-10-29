package kernel

import (
	"time"
)

type Schedule interface {
	// Id 任务标识，如果不提供将自动生成
	Id() uint64

	// Target 目标
	Target() uint64

	// Elapsed 预计消耗时间
	Elapsed() time.Duration

	// Type 类型
	Type() Type

	// Subtype 子类型
	Subtype() Type

	// Data 数据
	Data() map[string]any

	// Next 下一个可被执行的时间
	Next() time.Time
}
