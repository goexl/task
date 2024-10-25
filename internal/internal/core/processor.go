package core

// Processor 处理器
type Processor interface {
	// Process 处理任务调度
	Process(Tasking) (Executor, error)
}
