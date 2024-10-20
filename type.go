package task

import (
	"github.com/goexl/task/internal/core"
)

const (
	TypeCron Type = iota + 1
	TypeFixed
	TypeRate
	TypeComputable
	TypeOnce
)

type Type = core.Type
