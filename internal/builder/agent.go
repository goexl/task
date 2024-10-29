package builder

import (
	"github.com/goexl/log"
	"github.com/goexl/task/internal/core"
	"github.com/goexl/task/internal/param"
)

type Agent struct {
	params *param.Agent
}

func NewAgent(tasker core.Tasker) *Agent {
	return &Agent{
		params: param.NewAgent(tasker),
	}
}

func (a *Agent) Logger(logger log.Logger) (agent *Agent) {
	a.params.Logger = logger
	agent = a

	return
}

func (a *Agent) Retries(times uint32) (agent *Agent) {
	a.params.Retries = times
	agent = a

	return
}

func (a *Agent) Build() *core.Agent {
	return core.NewAgent(a.params)
}
