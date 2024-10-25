package builder

import (
	"time"

	"github.com/goexl/task/internal/internal"
	"github.com/goexl/task/internal/internal/constant"
	"github.com/goexl/task/internal/internal/param"
	"github.com/goexl/task/internal/kernel"
)

type Scheduling struct {
	params *param.Scheduling
}

func NewScheduling(target uint64, subtype kernel.Type) *Scheduling {
	return &Scheduling{
		params: param.NewScheduling(target, subtype),
	}
}

func (s *Scheduling) Id(id uint64) (scheduling *Scheduling) {
	s.params.Id = id
	scheduling = s

	return
}

func (s *Scheduling) Removable() (scheduling *Scheduling) {
	s.params.Removable = true
	scheduling = s

	return
}

func (s *Scheduling) Cron(expression string) (scheduling *Scheduling) {
	s.params.Type = kernel.TypeCron
	s.params.Data[constant.FieldCron] = expression
	scheduling = s

	return
}

func (s *Scheduling) Fixed(runtime time.Time) (scheduling *Scheduling) {
	s.params.Type = kernel.TypeFixed
	s.params.Data[constant.FieldRuntime] = runtime
	scheduling = s

	return
}

func (s *Scheduling) Rate(duration time.Duration) (scheduling *Scheduling) {
	s.params.Type = kernel.TypeRate
	s.params.Data[constant.FieldRate] = duration
	scheduling = s

	return
}

func (s *Scheduling) Computable() (scheduling *Scheduling) {
	s.params.Type = kernel.TypeComputable
	scheduling = s

	return
}

func (s *Scheduling) Once() (scheduling *Scheduling) {
	s.params.Type = kernel.TypeOnce
	scheduling = s

	return
}

func (s *Scheduling) Data(data any) (scheduling *Scheduling) {
	s.params.Data[constant.FieldData] = data
	scheduling = s

	return
}

func (s *Scheduling) Build() *internal.Scheduling {
	return internal.NewScheduling(s.params)
}
