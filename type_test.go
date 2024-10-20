package task

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestType(t *testing.T) {
	require.Equal(t, TypeCron, 1, "Cron任务定义出错")
	require.Equal(t, TypeFixed, 2, "固定时间任务定义出错")
	require.Equal(t, TypeRate, 3, "周期频率任务定义出错")
	require.Equal(t, TypeComputable, 4, "可计算任务定义出错")
	require.Equal(t, TypeOnce, 5, "只执行一次任务定义出错")
}
