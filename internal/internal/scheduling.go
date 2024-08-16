package internal

import (
	"github.com/goexl/task/internal/param"
)

type Scheduling struct {
	params *param.Scheduling
}

func NewScheduling(params *param.Scheduling) *Scheduling {
	return &Scheduling{
		params: params,
	}
}

func (s *Scheduling) Id() uint64 {
	return s.params.Id
}

func (s *Scheduling) Target() uint64 {
	return s.params.Target
}

func (s *Scheduling) Data() any {
	return s.params.Data
}
