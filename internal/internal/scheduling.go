package internal

import (
	"github.com/goexl/task/internal/internal/param"
	"github.com/goexl/task/internal/kernel"
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

func (s *Scheduling) Removable() bool {
	return s.params.Removable
}

func (s *Scheduling) Type() kernel.Type {
	return s.params.Type
}

func (s *Scheduling) Subtype() kernel.Type {
	return s.params.Subtype
}

func (s *Scheduling) Data() map[string]any {
	return s.params.Data
}
