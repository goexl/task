package task_test

import (
	"testing"

	"github.com/goexl/snowflake"
	"github.com/goexl/task"
	"github.com/stretchr/testify/require"
)

func TestScheduling(t *testing.T) {
	id := snowflake.New().Build().Parse(1)
	scheduling := task.NewSchedule(id, 1).Build()
	require.Equal(t, scheduling.Target(), id, "计划标识设置出错")
}
