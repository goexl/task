package task

// Processor 处理器
type Processor interface {
	// Process 处理任务调度
	Process(scheduling Scheduling) error
}
