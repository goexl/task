package kernel

import (
	"time"

	"github.com/goexl/id"
)

type Schedule interface {
	// Id 任务标识，如果不提供将自动生成
	Id() id.Value

	// Target 目标
	Target() id.Value

	// Timeout 任务超时时间
	Timeout() time.Duration

	// Type 类型
	Type() Type

	// Subtype 子类型
	Subtype() Type

	// Maximum 最大重试次数
	Maximum() uint32

	// Data 数据
	Data() map[string]any

	// Next 下一个可被执行的时间
	Next() time.Time
}
