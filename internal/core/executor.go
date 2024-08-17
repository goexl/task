package core

import (
	"time"
)

type Executor interface {
	// Execute 执行任务
	Execute() error

	// Recyclable 是否继续执行
	Recyclable() bool

	// Next 下次执行时间
	Next() time.Time

	// Error 执行错误
	Error() error
}
