package kernel

type Selector interface {
	// Select 选择不同的执行器执行任务
	Select(Task) (Executor, error)
}
