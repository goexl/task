package core

type Scheduling interface {
	// Id 任务标识，如果不提供将自动生成
	Id() uint64

	// Target 目标
	Target() uint64

	// Type 类型
	Type() Type

	// Data 数据
	Data() any
}
