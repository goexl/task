package core

type Tasking interface {
	Scheduling

	// Times 重试次数
	Times() uint32
}
