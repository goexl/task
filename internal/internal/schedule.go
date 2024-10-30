package internal

import (
	"time"

	"github.com/goexl/task/internal/internal/constant"
	"github.com/goexl/task/internal/internal/param"
	"github.com/goexl/task/internal/kernel"
	"github.com/robfig/cron/v3"
)

type Schedule struct {
	params *param.Schedule
	parser cron.Parser
}

func NewSchedule(params *param.Schedule) *Schedule {
	return &Schedule{
		params: params,
		parser: cron.NewParser(cron.Second),
	}
}

func (s *Schedule) Id() uint64 {
	return s.params.Id
}

func (s *Schedule) Target() uint64 {
	return s.params.Target
}

func (s *Schedule) Type() kernel.Type {
	return s.params.Type
}

func (s *Schedule) Subtype() kernel.Type {
	return s.params.Subtype
}

func (s *Schedule) Next() (next time.Time) {
	switch s.Type() {
	case kernel.TypeCron:
		next = s.nextCron()
	case kernel.TypeFixed:
		next = s.Data()[constant.FieldRuntime].(time.Time)
	case kernel.TypeRate:
		next = time.Now().Add(s.Data()[constant.FieldRate].(time.Duration))
	default:
		next = time.Now()
	}

	return
}

func (s *Schedule) Elapsed() time.Duration {
	return s.params.Elapsed
}

func (s *Schedule) Data() map[string]any {
	return s.params.Data
}

func (s *Schedule) nextCron() (next time.Time) {
	expression := s.Data()[constant.FieldCron].(string)
	if schedule, pe := s.parser.Parse(expression); nil != pe {
		next = time.Now()
	} else {
		next = schedule.Next(time.Now())
	}

	return
}
