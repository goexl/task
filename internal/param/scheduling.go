package param

type Scheduling struct {
	Id     uint64
	Target uint64
	Data   any
}

func NewScheduling(target uint64) *Scheduling {
	return &Scheduling{
		Target: target,
	}
}
