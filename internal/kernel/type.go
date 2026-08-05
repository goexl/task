package kernel

const (
	TypeUnknown    Type = 0
	TypeCron       Type = 1 // 表达式
	TypeFixed      Type = 2 // 固定时间
	TypeRate       Type = 3 // 固定频率
	TypeComputable Type = 4 // 自身计算
	TypeOnce       Type = 5 // 一次性任务
)

type Type uint16
