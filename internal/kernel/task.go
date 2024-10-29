package kernel

type Task interface {
	Schedule

	// Retries 重试次数
	Retries() uint32
}
