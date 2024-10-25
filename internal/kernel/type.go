package kernel

const (
	TypeCron Type = iota + 1
	TypeFixed
	TypeRate
	TypeComputable
	TypeOnce
)

type Type uint16
