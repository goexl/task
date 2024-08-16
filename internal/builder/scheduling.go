package builder

import (
	"github.com/goexl/task/internal/core"
	"github.com/goexl/task/internal/param"
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

func (s *Scheduling) Build() *core.Scheduling {
	return core.NewScheduling(s.params)
}
