package task_test

import (
	"testing"

	"github.com/goexl/task"
	"github.com/stretchr/testify/require"
)

func TestScheduling(t *testing.T) {
	scheduling := task.NewScheduling(1, 1).Build()
	require.Equal(t, scheduling.Target(), uint64(1), "计划标识设置出错")
}
