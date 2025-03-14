package kernel

type Task interface {
	Schedule

	// Times 当然运行次数
	Times() uint32
}
