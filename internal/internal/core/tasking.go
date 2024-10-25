package core

type Tasking interface {
	Scheduling

	// Retries 重试次数
	Retries() uint32
}
