package core

const (
	TypeCron Type = iota + 1
	TypeFixed
	TypeRate
	TypeComputable
	TypeOnce
)

type Type uint16

func (t Type) Cron() bool {
	return TypeCron == t
}

func (t Type) Fixed() bool {
	return TypeFixed == t
}

func (t Type) Rate() bool {
	return TypeRate == t
}

func (t Type) Computable() bool {
	return TypeComputable == t
}

func (t Type) Once() bool {
	return TypeOnce == t
}
