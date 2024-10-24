package task

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatus(t *testing.T) {
	require.Equal(t, StatusCreated, 1, "状态已创建定义出错")
	require.Equal(t, StatusRunning, 2, "状态运行中定义出错")
	require.Equal(t, StatusRetrying, 3, "状态重试中定义出错")

	require.Equal(t, StatusFailed, 20, "状态错误定义出错")

	require.Equal(t, StatusSuccess, 30, "状态成功定义出错")
}
