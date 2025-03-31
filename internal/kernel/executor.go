package kernel

type Executor interface {
	// Execute 执行任务
	Execute(ctx *Context, target uint64, times uint32) error
}
