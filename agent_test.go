package task_test

import (
	"testing"
	"time"

	"github.com/goexl/task"
	"github.com/goexl/task/internal/core"
	"github.com/stretchr/testify/require"
)

type tasker struct{}

func (t *tasker) Add(_ task.Scheduling) (err error) {
	return
}

func (t *tasker) Remove(_ core.Scheduling) (err error) {
	return
}

func (t *tasker) Running(_ uint64, _ task.Status, _ uint32) (err error) {
	return
}

func (t *tasker) Update(_ uint64, _ task.Status, _ time.Time) (err error) {
	return
}

func (t *tasker) Next(_ uint64) (err error) {
	return
}

func TestAgent(t *testing.T) {
	agent := task.New(new(tasker)).Build()
	require.NotNil(t, agent, "执行器创建出错")
}
