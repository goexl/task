package core

type Scheduling interface {
	// Id 任务标识，如果不提供将自动生成
	Id() uint64

	// Target 目标
	Target() uint64

	// Subtype 类型
	Type() Subtype

	// Data 数据
	Data() any
}
