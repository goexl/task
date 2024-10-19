package task

import (
	"github.com/goexl/task/internal/core"
)

const (
	TypeCron Type = iota + 1
	TypeFixed
	TypeRate
	TypeNext
)

// Type 类型
// 方便通过 Type 来定制化处理流程
type Type = core.Type
