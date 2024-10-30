package task_test

import (
	"testing"

	"github.com/goexl/task"
	"github.com/stretchr/testify/require"
)

func TestAgent(t *testing.T) {
	agent := task.New().Build()
	require.NotNil(t, agent, "执行器创建出错")
}
