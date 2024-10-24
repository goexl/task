package task

import (
	"time"
)

type Executor interface {
	// Execute 执行任务
	Execute() error

	// Next 下次执行时间
	Next() (bool, time.Time)

	// Error 执行错误
	Error() error
}
