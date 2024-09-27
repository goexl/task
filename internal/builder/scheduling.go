package builder

import (
	"github.com/goexl/task/internal/internal"
	"github.com/goexl/task/internal/internal/param"
)

type Scheduling struct {
	params *param.Scheduling
}

func NewScheduling(target uint64) *Scheduling {
	return &Scheduling{
		params: param.NewScheduling(target),
	}
}

func (s *Scheduling) Id(id uint64) (scheduling *Scheduling) {
	s.params.Id = id
	scheduling = s

	return
}

func (s *Scheduling) Data(data any) (scheduling *Scheduling) {
	s.params.Data = data
	scheduling = s

	return
}

func (s *Scheduling) Build() *internal.Scheduling {
	return internal.NewScheduling(s.params)
}
