package task

// Processor 处理器
type Processor[T any] interface {
	// Process 处理任务调度
	Process(scheduling T) error
}
