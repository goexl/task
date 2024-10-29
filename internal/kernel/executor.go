package kernel

type Executor interface {
	// Execute 执行任务
	Execute(target uint64, retries uint32) error

	// Error 执行错误
	Error() error
}
