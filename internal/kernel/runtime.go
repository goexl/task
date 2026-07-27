package kernel

import (
	"time"
)

type Runtime interface {
	// Id 任务标识，如果不提供将自动生成
	Id() uint64

	// Target 目标
	Target() uint64

	// Timeout 任务超时时间
	Timeout() time.Duration

	// Type 类型
	Type() Type

	// Subtype 子类型
	Subtype() Type

	// Maximum 最大重试次数
	Maximum() uint32

	// Next 下一个可被执行的时间
	Next() time.Time
}
