package task

// Agent 任务执行器
type Agent interface {
	// Start 开始接收任务
	Start(processor Processor) error

	// Add 添加
	Add(scheduling any) error

	// Remove 删除
	Remove(scheduling any) error
}
