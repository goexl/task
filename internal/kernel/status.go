package kernel

const (
	StatusCreated  = 1 // 已创建
	StatusRunning  = 2 // 运行中
	StatusRetrying = 3 // 重试中

	StatusSuccess = 10 // 成功
	StatusStandby = 11 // 待命

	StatusFailed = 255 // 已失败
)

type Status uint8
