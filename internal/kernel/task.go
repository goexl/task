package kernel

type Task interface {
	Runtime

	// Data 数据
	Data() any

	// Times 当然运行次数
	Times() uint32
}
