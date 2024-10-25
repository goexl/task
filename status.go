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

	// StatusFailed 失败
	StatusFailed = kernel.StatusFailed

	// StatusSuccess 成功
	StatusSuccess = kernel.StatusSuccess
)

// Status 类型
type Status = kernel.Status
