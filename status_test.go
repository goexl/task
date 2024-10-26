package task_test

import (
	"testing"

	"github.com/goexl/task"
	"github.com/stretchr/testify/require"
)

func TestStatus(t *testing.T) {
	require.Equal(t, task.StatusCreated, 1, "状态已创建定义出错")
	require.Equal(t, task.StatusRunning, 2, "状态运行中定义出错")
	require.Equal(t, task.StatusRetrying, 3, "状态重试中定义出错")

	require.Equal(t, task.StatusFailed, 10, "状态错误定义出错")

	require.Equal(t, task.StatusSuccess, 20, "状态成功定义出错")
}
