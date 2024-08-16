package internal

type Scheduling struct {
	target int64
	data   any
}

func NewScheduling(target int64, data any) *Scheduling {
	return &Scheduling{
		target: target,
		data:   data,
	}
}

func (s *Scheduling) Target() int64 {
	return s.target
}

func (s *Scheduling) Data() any {
	return s.data
}
