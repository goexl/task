package task

import (
	"github.com/goexl/task/internal/kernel"
)

const (
	// StatusCreated 已创建
	StatusCreated = kernel.StatusCreated
	// StatusRunning 执行中
	StatusRunning = kernel.StatusRunning
	// StatusRetrying 重试中
	StatusRetrying = kernel.StatusRetrying

	// StatusSuccess 成功
	StatusSuccess = kernel.StatusSuccess
	// StatusStandby 待命
	StatusStandby = kernel.StatusStandby

	// StatusFailed 失败
	StatusFailed = kernel.StatusFailed
)

// Status 类型
type Status = kernel.Status
