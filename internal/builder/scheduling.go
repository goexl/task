package builder

import (
	"time"

	"github.com/goexl/task/internal/core"
	"github.com/goexl/task/internal/internal"
	"github.com/goexl/task/internal/internal/constant"
	"github.com/goexl/task/internal/internal/param"
)

type Scheduling struct {
	params *param.Scheduling
}

func NewScheduling(target uint64, subtype core.Type) *Scheduling {
	return &Scheduling{
		params: param.NewScheduling(target, subtype),
	}
}

func (s *Scheduling) Id(id uint64) (scheduling *Scheduling) {
	s.params.Id = id
	scheduling = s

	return
}

func (s *Scheduling) Cron(expression string) (scheduling *Scheduling) {
	s.params.Type = core.TypeCron
	s.params.Data[constant.FieldCron] = expression
	scheduling = s

	return
}

func (s *Scheduling) Fixed(runtime time.Time) (scheduling *Scheduling) {
	s.params.Type = core.TypeFixed
	s.params.Data[constant.FieldRuntime] = runtime
	scheduling = s

	return
}

func (s *Scheduling) Rate(duration time.Duration) (scheduling *Scheduling) {
	s.params.Type = core.TypeRate
	s.params.Data[constant.FieldRate] = duration
	scheduling = s

	return
}

func (s *Scheduling) Computable() (scheduling *Scheduling) {
	s.params.Type = core.TypeComputable
	scheduling = s

	return
}

func (s *Scheduling) Once() (scheduling *Scheduling) {
	s.params.Type = core.TypeOnce
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
