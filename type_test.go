package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestType(t *testing.T) {
	assert.Equal(t, TypeCron, 1, "Cron任务定义出错")
	assert.Equal(t, TypeFixed, 2, "固定时间任务定义出错")
	assert.Equal(t, TypeRate, 3, "周期频率任务定义出错")
	assert.Equal(t, TypeNext, 4, "可计算任务定义出错")
}
