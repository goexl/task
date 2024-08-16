package task

// Agent 任务执行器
type Agent[T any] interface {
	// Start 开始接收任务
	Start(processor Processor[T]) error

	// Add 添加
	Add(scheduling T) error

	// Remove 删除
	Remove(scheduling T) error
}
