package builder

import (
	"time"

	"github.com/goexl/task/internal/internal"
	"github.com/goexl/task/internal/internal/constant"
	"github.com/goexl/task/internal/internal/param"
	"github.com/goexl/task/internal/kernel"
)

type Schedule struct {
	params *param.Schedule
}

func NewSchedule(target uint64, subtype kernel.Type) *Schedule {
	return &Schedule{
		params: param.NewSchedule(target, subtype),
	}
}

func (s *Schedule) Id(id uint64) (scheduling *Schedule) {
	s.params.Id = id
	scheduling = s

	return
}

func (s *Schedule) Cron(expression string) (scheduling *Schedule) {
	s.params.Type = kernel.TypeCron
	s.params.Data[constant.FieldCron] = expression
	scheduling = s

	return
}

func (s *Schedule) Fixed(runtime time.Time) (scheduling *Schedule) {
	s.params.Type = kernel.TypeFixed
	s.params.Data[constant.FieldRuntime] = runtime
	scheduling = s

	return
}

func (s *Schedule) Rate(duration time.Duration) (scheduling *Schedule) {
	s.params.Type = kernel.TypeRate
	s.params.Data[constant.FieldRate] = duration
	scheduling = s

	return
}

func (s *Schedule) Computable() (scheduling *Schedule) {
	s.params.Type = kernel.TypeComputable
	scheduling = s

	return
}

func (s *Schedule) Once() (scheduling *Schedule) {
	s.params.Type = kernel.TypeOnce
	scheduling = s

	return
}

func (s *Schedule) Data(data any) (scheduling *Schedule) {
	s.params.Data[constant.FieldData] = data
	scheduling = s

	return
}

func (s *Schedule) Build() *internal.Schedule {
	return internal.NewSchedule(s.params)
}
