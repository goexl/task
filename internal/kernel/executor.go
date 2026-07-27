package kernel

type Executor interface {
	// Execute 执行任务
	Execute(*Context, Task) error
}
